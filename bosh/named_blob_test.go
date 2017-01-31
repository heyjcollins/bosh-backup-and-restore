package bosh_test

import (
	"bytes"
	"fmt"
	"log"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	backuperfakes "github.com/pivotal-cf/pcf-backup-and-restore/backuper/fakes"
	"github.com/pivotal-cf/pcf-backup-and-restore/bosh"
	"github.com/pivotal-cf/pcf-backup-and-restore/bosh/fakes"
)

var _ = Describe("NamedBlob", func() {

	var sshConnection *fakes.FakeSSHConnection
	var boshLogger boshlog.Logger
	var instance *backuperfakes.FakeInstance
	var stdout, stderr *gbytes.Buffer
	var job bosh.Job

	var namedBlob *bosh.NamedBlob

	BeforeEach(func() {
		sshConnection = new(fakes.FakeSSHConnection)
		instance = new(backuperfakes.FakeInstance)
		job = bosh.NewJob(bosh.BackupAndRestoreScripts{"/var/vcap/jobs/foo1/p-backup"}, "named-blob")

		stdout = gbytes.NewBuffer()
		stderr = gbytes.NewBuffer()
		boshLogger = boshlog.New(boshlog.LevelDebug, log.New(stdout, "[bosh-package] ", log.Lshortfile), log.New(stderr, "[bosh-package] ", log.Lshortfile))

	})

	JustBeforeEach(func() {
		namedBlob = bosh.NewNamedBlob(instance, job, sshConnection, boshLogger)
	})

	Describe("StreamFromRemote", func() {
		var err error
		var writer = bytes.NewBufferString("dave")

		JustBeforeEach(func() {
			err = namedBlob.StreamFromRemote(writer)
		})

		Describe("when successful", func() {
			BeforeEach(func() {
				sshConnection.StreamReturns([]byte("not relevant"), 0, nil)
			})

			It("uses the ssh connection to tar the backup and stream it to the local machine", func() {
				Expect(sshConnection.StreamCallCount()).To(Equal(1))

				cmd, returnedWriter := sshConnection.StreamArgsForCall(0)
				Expect(cmd).To(Equal("sudo tar -C /var/vcap/store/backup/named-blob -zc ."))
				Expect(returnedWriter).To(Equal(writer))
			})

			It("does not fail", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Describe("when there is an error tarring the backup", func() {
			BeforeEach(func() {
				sshConnection.StreamReturns([]byte("not relevant"), 1, nil)
			})

			It("uses the ssh connection to tar the backup and stream it to the local machine", func() {
				Expect(sshConnection.StreamCallCount()).To(Equal(1))

				cmd, returnedWriter := sshConnection.StreamArgsForCall(0)
				Expect(cmd).To(Equal("sudo tar -C /var/vcap/store/backup/named-blob -zc ."))
				Expect(returnedWriter).To(Equal(writer))
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		Describe("when there is an SSH error", func() {
			var sshError error

			BeforeEach(func() {
				sshError = fmt.Errorf("I have the best SSH")
				sshConnection.StreamReturns([]byte("not relevant"), 0, sshError)
			})

			It("uses the ssh connection to tar the backup and stream it to the local machine", func() {
				Expect(sshConnection.StreamCallCount()).To(Equal(1))

				cmd, returnedWriter := sshConnection.StreamArgsForCall(0)
				Expect(cmd).To(Equal("sudo tar -C /var/vcap/store/backup/named-blob -zc ."))
				Expect(returnedWriter).To(Equal(writer))
			})

			It("fails", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(sshError))
			})
		})
	})

	Describe("Name", func() {
		It("returns the blob", func() {
			Expect(namedBlob.Name()).To(Equal("named-blob"))
		})
	})

	Describe("IsNamed", func() {
		It("returns true", func() {
			Expect(namedBlob.IsNamed()).To(BeTrue())
		})
	})

	Describe("BackupChecksum", func() {
		var actualChecksum map[string]string
		var actualChecksumError error

		JustBeforeEach(func() {
			actualChecksum, actualChecksumError = namedBlob.BackupChecksum()
		})

		Context("triggers find & shasum as root", func() {
			BeforeEach(func() {
				sshConnection.RunReturns([]byte("not relevant"), nil, 0, nil)
			})

			It("generates the correct request", func() {
				Expect(sshConnection.RunArgsForCall(0)).To(Equal("cd /var/vcap/store/backup/named-blob; sudo sh -c 'find . -type f | xargs shasum'"))
			})
		})
		Context("can calculate checksum", func() {
			BeforeEach(func() {
				sshConnection.RunReturns([]byte("07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e  file1\n07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e  file2\nn87fc29fb3aacd99f7f7b81df9c43b13e71c56a1e file3/file4"), nil, 0, nil)
			})
			It("converts the checksum to a map", func() {
				Expect(actualChecksumError).NotTo(HaveOccurred())
				Expect(actualChecksum).To(Equal(map[string]string{
					"file1":       "07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e",
					"file2":       "07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e",
					"file3/file4": "n87fc29fb3aacd99f7f7b81df9c43b13e71c56a1e",
				}))
			})
		})
		Context("can calculate checksum, with trailing spaces", func() {
			BeforeEach(func() {
				sshConnection.RunReturns([]byte("07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e file1\n"), nil, 0, nil)
			})
			It("converts the checksum to a map", func() {
				Expect(actualChecksumError).NotTo(HaveOccurred())
				Expect(actualChecksum).To(Equal(map[string]string{
					"file1": "07fc29fb3aacd99f7f7b81df9c43b13e71c56a1e",
				}))
			})
		})
		Context("sha output is empty", func() {
			BeforeEach(func() {
				sshConnection.RunReturns([]byte(""), nil, 0, nil)
			})
			It("converts an empty map", func() {
				Expect(actualChecksumError).NotTo(HaveOccurred())
				Expect(actualChecksum).To(Equal(map[string]string{}))
			})
		})
		Context("sha for a empty directory", func() {
			BeforeEach(func() {
				sshConnection.RunReturns([]byte("da39a3ee5e6b4b0d3255bfef95601890afd80709  -"), nil, 0, nil)
			})
			It("reject '-' as a filename", func() {
				Expect(actualChecksumError).NotTo(HaveOccurred())
				Expect(actualChecksum).To(Equal(map[string]string{}))
			})
		})

		Context("fails to calculate checksum", func() {
			expectedErr := fmt.Errorf("some error")

			BeforeEach(func() {
				sshConnection.RunReturns(nil, nil, 0, expectedErr)
			})
			It("returns an error", func() {
				Expect(actualChecksumError).To(MatchError(expectedErr))
			})
		})
		Context("fails to execute the command", func() {
			BeforeEach(func() {
				sshConnection.RunReturns(nil, nil, 1, nil)
			})
			It("returns an error", func() {
				Expect(actualChecksumError).To(HaveOccurred())
			})
		})
	})
})
