package migrations

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func MigrateUpProductsSuppliers(db *gorm.DB) {
	db.AutoMigrate(&models.ProductsSuppliers{})
	/*db.Exec(`
			CREATE TRIGGER add_stock_trigger
			AFTER INSERT ON products_suppliers
			FOR EACH ROW
			BEGIN
				UPDATE
					suppliers
				SET
					stock = stock + NEW.stock
				WHERE
					id = NEW.supplier_id;
				UPDATE
					products
				SET
					stock = stock + NEW.stock
				WHERE
					id = NEW.product_id;
			END;
		`)
	db.Exec(`
		CREATE TRIGGER remove_stock_trigger
		AFTER DELETE ON products_suppliers
		FOR EACH ROW
		BEGIN
			UPDATE
				suppliers
			SET
				stock = stock - OLD.stock
			WHERE
				id = OLD.supplier_id;
			UPDATE
				products
			SET
				stock = stock - OLD.stock
			WHERE
				id = OLD.product_id;
		END;
	`)
	db.Exec(`
		CREATE TRIGGER update_stock_trigger AFTER UPDATE ON products_suppliers
		FOR EACH ROW
		BEGIN
			UPDATE suppliers
			SET stock = stock - OLD.stock + NEW.stock
			WHERE id = OLD.supplier_id;

			UPDATE products
			SET stock = stock - OLD.stock + NEW.stock
			WHERE id = OLD.product_id;
		END;
	    `)*/
}

func MigrateDownProductsSuppliers(db *gorm.DB) {
	db.Migrator().DropTable("products_suppliers")
}
