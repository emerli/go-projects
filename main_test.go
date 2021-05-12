package main
//
// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"testing"
// 	"time"
//
// 	"delide-digidesk.module/delide-digidesk/delide-digidesk-api/connectionsfactory"
// 	"delide-digidesk.module/delide-digidesk/delide-digidesk-api/handlers"
// 	"delide-digidesk.module/delide-digidesk/delide-digidesk-api/models"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/gorm"
// )
//
// func TestBase(t *testing.T) {
// 	t.Run("alive", func(t *testing.T) {
// 		req := httptest.NewRequest(http.MethodGet, "/", nil)
//
// 		rr := httptest.NewRecorder()
// 		handler := handlers.IsAlive()
// 		handler.ServeHTTP(rr, req)
//
// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong s code: got %v want %v", status, http.StatusOK)
// 		}
//
// 		var r models.BaseResp
// 		if err := json.Unmarshal([]byte(rr.Body.String()), &r); err != nil {
// 			t.Errorf("parse error %s", err.Error())
// 		}
//
// 		assert.True(t, r.Success)
// 		assert.Contains(t, r.Data, "I'm alive")
// 	})
// }
//
// //
// // func TestConnections(t *testing.T) {
// // 	db, err := connectionsfactory.NewDB()
// // 	if err != nil {
// // 		log.Printf("exceptionsError while initializing context: %s\n", err.Error())
// // 	} else {
// // 		log.Printf("database connection Ok! ")
// // 	}
// //
// // 	router := CreateMuxRouter(db)
// //
// // 	req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// // 	rr := httptest.NewRecorder()
// //
// // 	router.ServeHTTP(rr, req)
// //
// //
// //
// //
// // 	t. Run("conns", func(t *testing.T) {
// // 		if status := rr.Code; status != http.StatusOK {
// // 			t.Errorf("handler returned wrong s code: got %v want %v", status, http.StatusOK)
// // 		}
// //
// // 		var r models.BaseResp
// // 		if err := json.Unmarshal([]byte(rr.Body.String()), &r); err != nil {
// // 			t.Errorf("parse error %s", err.Error())
// // 		}
// //
// // 		assert.True(t, r.Success)
// // 		assert.Contains(t, r.Data, "I'm alive")
// // 	})
// // }
//
// func BenchmarkTemplateParallel(b *testing.B) {
// 	// db, err := connectionsfactory.NewDB()
// 	// if err != nil {
// 	// 	log.Printf("exceptionsError while initializing context: %s\n", err.Error())
// 	// } else {
// 	log.Println("database connection Ok! ")
// 	// }
//
// 	// router := CreateMuxRouter(db)
//
// 	i := 10
// 	b.Run(strconv.Itoa(i), func(b *testing.B) {
// 		log.Println("database connection Ok! ")
// 		b.SetParallelism(i)
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
//
// 				//				log.Println("starting test ")
// 				// req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// 				// rr := httptest.NewRecorder()
// 				//
// 				// router.ServeHTTP(rr, req)
// 				//
// 				// if status := rr.Code; status != http.StatusOK {
// 				// 	log.Printf("bad status! ")
// 				// }
// 			}
// 		})
// 	})
// 	//
// 	//
// 	// 	b.SetParallelism(10)
// 	// 	b.RunParallel(func(pb *testing.PB) {
// 	// 		for  pb.Next() {
// 	//
// 	log.Println("starting test ")
// 	// 			req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// 	// 			rr := httptest.NewRecorder()
// 	//
// 	// 			router.ServeHTTP(rr, req)
// 	//
// 	// 			if status := rr.Code; status != http.StatusOK {
// 	// 				log.Printf("bad status! ")
// 	// 			}
// 	//
// 	// 		}
// 	//
// 	// 	})
// 	//
// 	// 	log.Println("waiting ")
// 	//
// 	// 	time.Sleep(time.Minute * 3)
// 	// 	log.Println("done")
// }
//
// func TestTeardownParallel(t *testing.T) {
// 	db, err := connectionsfactory.NewDB()
// 	if err != nil {
// 		log.Printf("exceptionsError while initializing context: %s\n", err.Error())
// 	} else {
// 		log.Println("database connection Ok! ")
// 	}
//
// 	router := CreateMuxRouter(db)
//
// 	t.Run("parallel", func(t *testing.T) {
// 		for i := 0; i < 200; i++ {
// 			t.Run("Test-"+strconv.Itoa(i), testConn(db, router))
// 		}
// 	})
//
// 	var c int
// 	db.Table("pg_stat_activity").Where("datname='ader'").Select("count(datid)").Scan(&c)
// 	log.Println(c)
//
// 	time.Sleep(time.Second * 90)
//
// 	db.Table("pg_stat_activity").Where("datname='ader'").Select("count(datid)").Scan(&c)
// 	log.Println(c)
//
// }
//
// func testConn(db *gorm.DB, router *mux.Router) func(t *testing.T) {
// 	return func(t *testing.T) {
// 		t.Parallel()
// 		log.Println("request " + t.Name())
// 		req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// 		rr := httptest.NewRecorder()
//
// 		router.ServeHTTP(rr, req)
//
// 		log.Println("response " + t.Name() + " got " + strconv.Itoa(rr.Code))
// 	}
// }
//
// func BenchmarkConns(b *testing.B) {
// 	db, err := connectionsfactory.NewDB()
// 	if err != nil {
// 		log.Printf("exceptionsError while initializing context: %s\n", err.Error())
// 	} else {
// 		log.Println("database connection Ok! ")
// 	}
//
// 	router := CreateMuxRouter(db)
//
//
//
// 	for x := 0; x < 1; x++ {
// 		b.SetParallelism(4)
// 		b.RunParallel(func(pb *testing.PB) {
// 			for pb.Next() {
// 				log.Println("request " + b.Name() )
// 				req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// 				rr := httptest.NewRecorder()
//
// 				router.ServeHTTP(rr, req)
//
// 				log.Println("response " + b.Name() + " got " + strconv.Itoa(rr.Code))
//
// 			}
// 		})
//
//
// 		// b.Run("test-"+strconv.Itoa(x), func(b *testing.B) {
// 		// 	for i := 0; i < b.N; i++ {
// 		// 		log.Println("request " + b.Name() + strconv.Itoa(i))
// 		// 		req := httptest.NewRequest(http.MethodGet, "/appointment/1", nil)
// 		// 		rr := httptest.NewRecorder()
// 		//
// 		// 		router.ServeHTTP(rr, req)
// 		//
// 		// 		log.Println("response " + b.Name() + strconv.Itoa(i)+ " got " + strconv.Itoa(rr.Code))
// 		// 	}
// 		// })
// 	}
//
// 	var c int
// 	db.Table("pg_stat_activity").Where("datname='ader'").Select("count(datid)").Scan(&c)
// 	log.Println(c)
//
// }
