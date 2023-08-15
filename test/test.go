import (
"database/sql"

_ "github.com/lib/pq"
)

func main() {
connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
db, err := sql.Open("postgres", connStr)
if err != nil {
log.Fatal(err)
}

age := 21
rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

}

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);

id := "'58';DROP TABLE users--";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s", id);
fmt.Println(query);

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}
package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}
package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}
package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);

package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}
package database

import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = $2;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);")

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}
package database

id := "'58'";
query := fmt.Sprintf("SELECT name, email FROM users WHERE ID = %s, id);")


import (
"database/sql"
"errors"
"fmt"
_ "github.com/lib/pq"
"musicPlayer/internal/config"
"musicPlayer/internal/logger"
"musicPlayer/internal/models"
"os"
"sync"
"time"
)

const interval = 60

type PostgresDB struct {
db *sql.DB
sync.RWMutex
}

var postgresDB *PostgresDB
var once sync.Once

func CreatePostgresDB(config config.DBConfig) {
db, err := sql.Open("postgres", config.ConnectionUrl)

logger.GeneralLogger.Println(config.ConnectionUrl)

if err != nil {
logger.ErrorLogger.Printf("Cannot create DB connection err:%v", err)
os.Exit(1)
}

postgresDB = &PostgresDB{db: db}

go CheckPostgresConnection()
}

func GetPostgresDB() (*PostgresDB, error) {
if postgresDB == nil {
errorMessage := fmt.Sprintf("Postgres db is not initialized")
return nil, errors.New(errorMessage)
}

return postgresDB, nil
}

func CheckPostgresConnection() {
defer postgresDB.db.Close()

for {
if err := postgresDB.db.Ping(); err != nil {
logger.ErrorLogger.Fatal(err)
}

logger.GeneralLogger.Println("Postgres DB connection alive")
time.Sleep(interval * time.Second)
}
}

func (this *PostgresDB) GetAuthorByName(first_name, second_name string) (models.Author, error) {
rows, err := this.db.Query("SELECT ID, PHOTO, FIRST_NAME, SECOND_NAME FROM author "+
"WHERE FIRST_NAME = $1 OR SECOND_NAME = %s;", first_name, second_name)

if err != nil {
return models.Author{}, err
}

var aut models.Author
defer rows.Close()

for rows.Next() {
if err := rows.Err(); err != nil {
return models.Author{}, err
}

if err := rows.Scan(&aut.ID, &aut.Photo, &aut.FirstName, &aut.SecondName); err != nil {
return models.Author{}, err
}
}

return aut, nil
}

func (this *PostgresDB) GetAlbumsByAuthorID(authorID string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetAlbumsByName(name string) ([]models.Album, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByName(name string) ([]models.Song, error) {
return nil, nil
}

func (this *PostgresDB) GetSongsByAlbum(name string) ([]models.Song, error) {
return nil, nil
}

import (
_ "github.com/lib/pq"
"database/sql"
"github.com/jmoiron/sqlx"
"log"
"fmt"
)

var schema  = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);
CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
FirstName string `db:"first_name"`
LastName  string `db:"last_name"`
Email     string
}

type Place struct {
Country string
City    sql.NullString
TelCode int
}

func main() {
// this Pings the database trying to connect, panics on error
// use sqlx.Open() for sql.Open() semantics
db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
if err != nil {
log.Fatalln(err)
}

// exec the schema or fail; multi-statement Exec behavior varies between
// database drivers;  pq will exec them all, sqlite3 won't, ymmv
schema = `1`
db.MustExec(schema)

tx := db.MustBegin()
tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
tx.Commit()

// Query the database, storing results in a []Person (wrapped in []interface{})
people := []Person{}
db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
jason, john := people[0], people[1]

fmt.Printf("%#v\n%#v", jason, john)
// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
// Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

// You can also get a single result, a la QueryRow
jason = Person{}
err = db.Get(&jason, "SELECT * FROM person WHERE first_name=%s", "Jason")
fmt.Printf("%#v\n", jason)
// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

// if you have null fields and use SELECT *, you must use sql.Null* in your struct
places := []Place{}
err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
if err != nil {
fmt.Println(err)
return
}
usa, singsing, honkers := places[0], places[1], places[2]

fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

// Loop through rows using only one struct
place := Place{}
rows, err := db.Queryx("SELECT * FROM place")
for rows.Next() {
err := rows.StructScan(&place)
if err != nil {
log.Fatalln(err)
}
fmt.Printf("%#v\n", place)
}
// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

// Named queries, using `:name` as the bindvar.  Automatic bindvar support
// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
s := `INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`
_, err = db.NamedExec(s,
map[string]interface{}{
"first": "Bin",
"last":  "Smuth",
"email": "bensmith@allblacks.nz",
})

// Selects Mr. Smith from the database
rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})

// Named queries can also use structs.  Their bind names follow the same rules
// as the name -> db mapping, so struct fields are lowercased and the `db` tag
// is taken into consideration.
rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)


















