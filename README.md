# GoApp

İlk proje dosyasını açtıktan sonra proje ismi ile init yapıyoruz

go mod init GoApp

daha sonra bir enviroment mekanizması oluşturabilmek için 

go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon <- Bunu neden yaptı anlamadım.


daha sonra;

go get github.com/joho/godotenv

Daha sonrasında Gin Frameworkünü kuruyoruz;

go get -u github.com/gin-gonic/gin

Veritabanı işlemleri için kullanacağımız orm;

go get -u gorm.io/gorm

Daha sonra veritabanı driverimizi kuruyoruz;

go get -u gorm.io/driver/postgres <- Biz postgre kullanacağımız için bunu yüklüyoruz.

Daha sonra aşağıda ki işlemi yapıyoruz

CompileDaemon -command="./GoApp" <- Daha fazla bilgi; https://github.com/githubnemo/CompileDaemon



projeyi başlatmak için; 

 CompileDaemon -command="./GoApp"

kütüphaneler;
https://gorm.io/docs/
https://gin-gonic.com/docs/quickstart/
https://github.com/joho/godotenv
https://www.elephantsql.com/ <- PC de postgresql varsa oda kullanılabilir.
