package models_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"jennifer/dealls-tech-test/internal/domain/models"
	"jennifer/dealls-tech-test/internal/server"
)

var _ = Describe("Migrate", func() {
	It("should be successful when automatically migrate database", func() {
		type DBModel struct {
			gorm.Model
			Name string
		}
		models.ModelLists = []interface{}{&DBModel{}}

		db, err := server.Conn(context.Background())

		Expect(err).NotTo(HaveOccurred())
		Expect(models.AutoMigrate(db)).To(Succeed())

		result := db.Create(&DBModel{
			Name: "test",
		})

		Expect(result.Error).NotTo(HaveOccurred())
		Expect(result.RowsAffected).To(BeEquivalentTo(1))
	})
})
