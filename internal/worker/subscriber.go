package worker

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"jubeliotesting/internal/dto"
	"jubeliotesting/internal/excel"
	"jubeliotesting/internal/service"
	"jubeliotesting/pkg/config"
	"jubeliotesting/pkg/manipulation"
	"jubeliotesting/pkg/parser"
	"log"
	"time"
)

type Subscriber struct {
	SalesService         *service.SalesService
	ProductService       *service.ProductService
	CategoryService      *service.CategoryService
	SubCategoryService   *service.SubCategoryService
	InventoryMoveService *service.InventoryMoveService
}

func NewSubcriber(
	salesService *service.SalesService,
	productService *service.ProductService,
	categoryService *service.CategoryService,
	subCategoryService *service.SubCategoryService,
	inventoryMoveService *service.InventoryMoveService,
) *Subscriber {
	return &Subscriber{
		SalesService:         salesService,
		ProductService:       productService,
		CategoryService:      categoryService,
		SubCategoryService:   subCategoryService,
		InventoryMoveService: inventoryMoveService,
	}
}

func (s *Subscriber) StartSubscriber(ctx context.Context, rdb *redis.Client, channel string, configsEnv config.GetEnvConfig) {
	sub := rdb.Subscribe(ctx, channel)
	ch := sub.Channel()

	for {
		select {
		case msg := <-ch:
			log.Println("Process Dataset Status:", msg.Payload)
			var publishDto dto.PublishDto
			if err := json.Unmarshal([]byte(msg.Payload), &publishDto); err != nil {
				log.Println(err)
				return
			}

			if publishDto.Key != configsEnv.APIKey {
				log.Println(errors.New("wrong API key"))
				return
			}

			excelReader := &excel.ExcelReaderImpl{}
			rows, err := excelReader.ReadExcel("dataset/dataset.xlsx")
			if err != nil {
				log.Println(err)
			}

			var categoryDTOArr []dto.CategoryDTO
			var subCategoryDTOArr []dto.SubCategoryDTO
			var supermarketDtoArr []dto.SupermarketDTO
			var productDTOArr []dto.ProductDTO

			for _, row := range rows {
				category := row[0]
				city := row[1]
				contry := row[2]
				customerName := row[3]
				manufacturer := row[4]
				orderDate, errod := parser.ParseDate(row[5])
				if errod != nil {
					log.Println(errod)
				}

				orderId := row[6]
				postalCode := row[7]
				productName := row[8]
				region := row[9]
				segment := row[10]
				shipDate, errshipdate := parser.ParseDate(row[11])
				if errshipdate != nil {
					log.Println(errshipdate)
				}

				shipMode := row[12]
				state := row[13]
				subCategory := row[14]
				discount, errdisc := parser.ParsePercentToFloat(row[15])
				if errdisc != nil {
					log.Println(errdisc)
				}

				numberOfRecords, errnor := parser.ParseStringInt(row[16])
				if errnor != nil {
					log.Println(errnor)
				}

				profit, errprof := parser.ParseDollarToFloat(row[17])
				if errprof != nil {
					log.Println(errprof)
				}

				profitRatio, errprofR := parser.ParsePercentToFloat(row[18])
				if errprofR != nil {
					log.Println(errprofR)
				}

				quantity, errqty := parser.ParseStringInt(row[19])
				if errqty != nil {
					log.Println(errqty)
				}

				sales, errsal := parser.ParseDollarToFloat(row[20])
				if errsal != nil {
					log.Println(errsal)
				}

				categoryDTOArr = manipulation.AddUniqueCategory(categoryDTOArr, category)
				subCategoryDTOArr = manipulation.AddUniqueSubCategory(subCategoryDTOArr, subCategory, category)
				productDTOArr = manipulation.AddUniqueProduct(productDTOArr, productName, subCategory, manufacturer)

				var supermarketDto dto.SupermarketDTO
				supermarketDto.Category = category
				supermarketDto.City = city
				supermarketDto.Country = contry
				supermarketDto.CustomerName = customerName
				supermarketDto.Manufacturer = manufacturer
				supermarketDto.OrderDate = orderDate
				supermarketDto.OrderID = orderId
				supermarketDto.PostalCode = postalCode
				supermarketDto.ProductName = productName
				supermarketDto.Region = region
				supermarketDto.Segment = segment
				supermarketDto.ShipDate = shipDate
				supermarketDto.ShipMode = shipMode
				supermarketDto.State = state
				supermarketDto.SubCategory = subCategory
				supermarketDto.Discount = discount
				supermarketDto.NumberOfRecord = numberOfRecords
				supermarketDto.Profit = profit
				supermarketDto.ProfitRatio = profitRatio
				supermarketDto.Quantity = quantity
				supermarketDto.Sales = sales

				supermarketDtoArr = append(supermarketDtoArr, supermarketDto)

				//fmt.Println(category, "|", city, "|", contry, "|", customerName, "|", manufacturer, "|", orderDate, "|", orderId, "|", postalCode, "|", productName, "|", region, "|", segment, "|", shipDate, "|", shipMode, "|", state, "|", subCategory, "|", discount, "|", numberOfRecords, "|", profit, "|", profitRatio, "|", quantity, "|", sales)
			}

			if publishDto.MasterData {
				for i, category := range categoryDTOArr {
					if err = s.CategoryService.CreateCategory(category); err != nil {
						log.Println(err)
					}
					percent := float64(i+1) / float64(len(supermarketDtoArr)) * 100
					key := publishDto.ClientID + ":process"
					rdb.Set(ctx, key, percent, time.Duration(2000)*time.Millisecond)
				}

				for i, subCategory := range subCategoryDTOArr {
					category, errget := s.CategoryService.GetCategoryByName(subCategory.CategoryName)
					if errget != nil {
						log.Println(errget)
					}

					subCategory.CategoryID = category.ID
					if err = s.SubCategoryService.CreateSubCategory(subCategory); err != nil {
						log.Println(err)
					}

					percent := float64(i+1) / float64(len(supermarketDtoArr)) * 100
					key := publishDto.ClientID + ":process"
					rdb.Set(ctx, key, percent, time.Duration(2000)*time.Millisecond)
				}

				for i, product := range productDTOArr {
					subcategory, errget := s.SubCategoryService.GetSubCategoryByName(product.SubCategoryName)
					if errget != nil {
						log.Println(errget)
					}

					product.SubCategoryID = subcategory.ID
					if err = s.ProductService.CreateProduct(product); err != nil {
						log.Println(err)
					}

					percent := float64(i+1) / float64(len(supermarketDtoArr)) * 100
					key := publishDto.ClientID + ":process"
					rdb.Set(ctx, key, percent, time.Duration(2000)*time.Millisecond)
				}
			}

			s.SalesService.TruncateSales()
			s.InventoryMoveService.TruncateInventory()

			for i, sale := range supermarketDtoArr {
				var saleDTO dto.SaleDTO

				category, errget := s.CategoryService.GetCategoryByName(sale.Category)
				if errget != nil {
					log.Println(errget)
				}

				subcategory, errgetsub := s.SubCategoryService.GetSubCategoryByName(sale.SubCategory)
				if errgetsub != nil {
					log.Println(errgetsub)
				}

				product, errgetprod := s.ProductService.GetProductsByName(sale.ProductName)
				if errgetprod != nil {
					log.Println(errgetprod)
				}

				saleDTO.CategoryID = category.ID
				saleDTO.City = sale.City
				saleDTO.Country = sale.Country
				saleDTO.CustomerName = sale.CustomerName
				saleDTO.Manufacturer = sale.Manufacturer
				saleDTO.OrderDate = sale.OrderDate
				saleDTO.OrderID = sale.OrderID
				saleDTO.PostalCode = sale.PostalCode
				saleDTO.ProductID = product.ID
				saleDTO.Region = sale.Region
				saleDTO.Segment = sale.Segment
				saleDTO.ShipDate = sale.ShipDate
				saleDTO.ShipMode = sale.ShipMode
				saleDTO.State = sale.State
				saleDTO.SubCategoryID = subcategory.ID
				saleDTO.Discount = sale.Discount
				saleDTO.NumberOfRecord = sale.NumberOfRecord
				saleDTO.Profit = sale.Profit
				saleDTO.ProfitRatio = sale.ProfitRatio
				saleDTO.Quantity = sale.Quantity
				saleDTO.Sales = sale.Sales

				if err = s.SalesService.Create(saleDTO); err != nil {
					log.Println(err)
				}

				var inventoryMoveDto dto.InventoryMoveDTO
				inventoryMoveDto.ProductID = product.ID
				inventoryMoveDto.OrderID = sale.OrderID
				inventoryMoveDto.Quantity = sale.Quantity
				inventoryMoveDto.OrderDate = sale.OrderDate
				inventoryMoveDto.MovementAt = sale.ShipDate
				inventoryMoveDto.Note = "-"

				if err = s.InventoryMoveService.CreateInventoryMove(inventoryMoveDto); err != nil {
					log.Println(err)
				}

				percent := float64(i+1) / float64(len(supermarketDtoArr)) * 100
				key := publishDto.ClientID + ":process"
				rdb.Set(ctx, key, percent, time.Duration(2000)*time.Millisecond)
				//fmt.Printf("Progress: %.2f%%\n", percent)
			}

			log.Println("Process Done!")
		case <-ctx.Done():
			log.Println("Subscriber shutting down...")
			return
		}
	}
}
