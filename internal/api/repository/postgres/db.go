package postgres

//
//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/google/uuid"
//	_ "github.com/lib/pq"
//	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/repository"
//	"log"
//	"strings"
//
//	"github.com/mavissig/GUC.DynamicPricing-API/internal/api/domain"
//)
//
//type Postgres struct {
//	DB *sql.DB
//}
//
//func New(cfg *repository.Config) *Postgres {
//	dsn := fmt.Sprintf("host=%s dbname=%s port=%v sslmode=disable",
//		cfg.PG.Host,
//		cfg.PG.Name,
//		cfg.PG.Port,
//	)
//
//	db, err := sql.Open("postgres", dsn)
//	if err != nil {
//		log.Fatalf("[DB][POSTGRES][ERROR]: create db : %v\n", err)
//	}
//	return &Postgres{
//		DB: db,
//	}
//}
//
//func (r *Postgres) InitTable(clearStart bool) {
//	if clearStart {
//		_, err := r.DB.Query(sqlResetTables)
//		if err != nil {
//			log.Printf("[DB][POSTGRES][ERROR]: reset tables: %v\n", err)
//		}
//	}
//	_, err := r.DB.Query(sqlInitTables)
//	if err != nil {
//		log.Printf("[DB][POSTGRES][ERROR]: initialization tables: %v\n", err)
//	}
//
//}
//
///*
//----------------------------------------------------------------
//                             CLIENT
//----------------------------------------------------------------
//*/
//
//func (r *Postgres) ClientAdd(client *domain.Client) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlAddClient, client.ClientName, client.ClientSurname, client.Birthday.Format("2006-01-02"), client.Gender, client.RegistrationDate.Format("2006-01-02"), client.AddressId))
//	return err
//}
//
//func (r *Postgres) ClientDeleteById(id uint64) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlDeleteById, "Client", id))
//	return err
//}
//
//func (r *Postgres) ClientGet(name, surname string, page, pageSize int) ([]*domain.Client, error) {
//	sqlReq := strings.Builder{}
//
//	sqlReq.WriteString("Client")
//	if name != "" || surname != "" {
//		sqlReq.WriteString(" WHERE ")
//		if name != "" {
//			sqlReq.WriteString(fmt.Sprintf(" client_name = '%s' ", name))
//		}
//		if name != "" && surname != "" {
//			sqlReq.WriteString(" AND ")
//		}
//		if surname != "" {
//			sqlReq.WriteString(fmt.Sprintf(" client_surname = '%s' ", surname))
//		}
//	}
//
//	row, err := r.DB.Query(fmt.Sprintf(sqlCommonGet, sqlReq.String()) + setPagination(page, pageSize))
//	if err != nil {
//		log.Printf("[DB][POSTGRES][ERROR]: get client by id: %v\n", err)
//		return nil, err
//	}
//
//	var clients []*domain.Client
//	for row.Next() {
//		client, err := parseClient(row)
//		if err != nil {
//			log.Printf("[DB][POSTGRES][ERROR]: get client by id error scan: %v\n", err)
//			return nil, err
//		}
//		clients = append(clients, client)
//	}
//
//	if len(clients) == 0 {
//		return nil, errors.New("not found")
//	}
//	return clients, nil
//}
//
//func (r *Postgres) ClientGetAll() ([]*domain.Client, error) {
//	row, err := r.DB.Query(fmt.Sprintf(sqlCommonGet, "Client"))
//	if err != nil {
//		log.Printf("[DB][POSTGRES][ERROR]: get client by id: %v\n", err)
//		return nil, err
//	}
//
//	var clients []domain.Client
//
//	for row.Next() {
//		client, err := parseClient(row)
//		if err != nil {
//			log.Printf("[DB][POSTGRES][ERROR]: get client by id error scan: %v\n", err)
//			return nil, err
//		}
//		clients = append(clients, *client)
//	}
//
//	fmt.Println("res:")
//	fmt.Println(clients)
//
//	return nil, nil
//}
//
//func (r *Postgres) ClientChangeAddress(id uint64, address *domain.Address) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlCommonChangeAddress, "Client"), address.Country, address.City, address.Street, id)
//	return err
//}
//
///*
//----------------------------------------------------------------
//                             PRODUCT
//----------------------------------------------------------------
//*/
//
//func (r *Postgres) ProductAdd(product *domain.Product) error {
//	imageId := "NULL"
//	if product.ImageId.String() != "00000000-0000-0000-0000-000000000000" {
//		imageId = fmt.Sprintf("'%s'", product.ImageId.String())
//	}
//
//	log.Println(imageId)
//
//	sqlReq := fmt.Sprintf(sqlAddProduct, product.Name, product.Category, product.Price, product.AvailableStock, product.LastUpdateDate.Format("2006-01-02"), product.SupplierId, imageId)
//
//	_, err := r.DB.Exec(sqlReq)
//	return err
//}
//
//func (r *Postgres) ProductGet(id uint64, page, pageSize int) ([]*domain.Product, error) {
//	var row *sql.Rows
//	var err error
//
//	if id != 0 {
//		row, err = r.DB.Query(fmt.Sprintf(sqlGetProduct, fmt.Sprintf("WHERE id = %d", id)))
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		row, err = r.DB.Query(fmt.Sprintf(sqlGetProduct, "") + setPagination(page, pageSize))
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	products := []*domain.Product{}
//	for row.Next() {
//		product, err := parseProduct(row)
//		if err != nil {
//			return nil, err
//		}
//		products = append(products, product)
//	}
//
//	if id != 0 && len(products) == 0 {
//		return products, errors.New("not found")
//	}
//
//	return products, nil
//}
//
//func (r *Postgres) ProductDelete(id uint64) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlDeleteById, "Product", id))
//	return err
//}
//
//func (r *Postgres) ProductDecrement(id, n uint64) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlDecProduct, id, n))
//	return err
//}
//
///*
//----------------------------------------------------------------
//                             SUPPLIER
//----------------------------------------------------------------
//*/
//
//func (r *Postgres) SupplierAdd(supplier *domain.Supplier) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlAddSupplier, supplier.Name, supplier.AddressId, supplier.PhoneNumber))
//	return err
//}
//
//func (r *Postgres) SupplierGet(id uint64, page, pageSize int) ([]*domain.Supplier, error) {
//	var row *sql.Rows
//	var err error
//
//	if id != 0 {
//		row, err = r.DB.Query(fmt.Sprintf(sqlCommonGetById, "Supplier", id))
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		row, err = r.DB.Query(fmt.Sprintf(sqlCommonGet, "Supplier") + setPagination(page, pageSize))
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	suppliers := []*domain.Supplier{}
//	for row.Next() {
//		supplier, err := parseSupplier(row)
//		if err != nil {
//			return nil, err
//		}
//		suppliers = append(suppliers, supplier)
//	}
//
//	if id != 0 && len(suppliers) == 0 {
//		return suppliers, errors.New("not found")
//	}
//
//	return suppliers, nil
//}
//
//func (r *Postgres) SupplierDelete(id uint64) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlDeleteById, "Supplier", id))
//	return err
//}
//
//func (r *Postgres) SupplierChangeAddress(id uint64, address *domain.Address) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlCommonChangeAddress, "Supplier"), address.Country, address.City, address.Street, id)
//	return err
//}
//
///*
//----------------------------------------------------------------
//                             IMAGES
//----------------------------------------------------------------
//*/
//
//func (r *Postgres) ImagesAdd(id uint64, b []byte) error {
//	_, err := r.DB.Exec(sqlAddImage, id, string(b))
//	return err
//}
//
//func (r *Postgres) ImagesGet(id uuid.UUID) ([]byte, error) {
//	row, err := r.DB.Query(fmt.Sprintf(sqlGetImageById, id.String()))
//	if err != nil {
//		return nil, err
//	}
//
//	row.Next()
//	img, err := parseImage(row)
//	if err != nil {
//		return nil, errors.New("not found")
//	}
//
//	return img, nil
//}
//
//func (r *Postgres) ImagesGetByProduct(id uint64) ([]byte, error) {
//	row, err := r.DB.Query(fmt.Sprintf(sqlGetImageByProductId, id))
//	if err != nil {
//		return nil, err
//	}
//
//	row.Next()
//	img, err := parseImage(row)
//	if err != nil {
//		return nil, errors.New("not found")
//	}
//
//	return img, nil
//}
//
//func (r *Postgres) ImagesDelete(id uuid.UUID) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlDeleteImage, id.String()))
//	return err
//}
//
//func (r *Postgres) ImagesChange(id uuid.UUID, b []byte) error {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlChangeImageById, id.String()), string(b))
//	return err
//}
//
///*
// ----------------------------------------------------------------
//                                 COMMON
// ----------------------------------------------------------------
//*/
//
//func parseClient(row *sql.Rows) (*domain.Client, error) {
//	res := &domain.Client{}
//	err := row.Scan(&res.Id, &res.ClientName, &res.ClientSurname, &res.Birthday, &res.Gender, &res.RegistrationDate, &res.AddressId)
//	return res, err
//}
//
//func parseProduct(row *sql.Rows) (*domain.Product, error) {
//	res := &domain.Product{}
//	err := row.Scan(&res.Id, &res.Name, &res.Category, &res.Price, &res.AvailableStock, &res.LastUpdateDate, &res.SupplierId, &res.ImageId)
//	return res, err
//}
//
//func parseSupplier(row *sql.Rows) (*domain.Supplier, error) {
//	res := &domain.Supplier{}
//	err := row.Scan(&res.Id, &res.Name, &res.AddressId, &res.PhoneNumber)
//	return res, err
//}
//
//func parseImage(row *sql.Rows) ([]byte, error) {
//	var res []byte
//	err := row.Scan(&res)
//	return res, err
//}
//
//func getOffset(page, pageSize int) int {
//	return (page - 1) * pageSize
//}
//
//func setPagination(page, pageSize int) string {
//	if pageSize < 1 {
//		return ""
//	}
//	return fmt.Sprintf(" LIMIT %d OFFSET %d ", pageSize, getOffset(page, pageSize))
//}
//
///*
// ----------------------------------------------------------------
//                           DEBUG METHODS
// ----------------------------------------------------------------
//*/
//
//func (r *Postgres) TestCreate() {
//	_, err := r.DB.Query(sqlDebugClient)
//	if err != nil {
//		log.Printf("[DB][POSTGRES][ERROR]: create debug client: %v\n", err)
//	}
//}
//
//func (r *Postgres) CommonChangeAddress(table, country, city, street string, tableId int) {
//	_, err := r.DB.Exec(fmt.Sprintf(sqlCommonChangeAddress, table), country, city, street, tableId)
//	if err != nil {
//		log.Printf("[DB][POSTGRES][ERROR]: CommonChangeAddress: %v\n", err)
//	}
//}
