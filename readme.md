Для примера, сервер запущен на порте :8888
Изменить порт можно по пути internal/config/local.yaml

Доступны 4 роута:
GET	http://localhost:8001/accounts/{id}/balance
		
POST	http://localhost:8001/accounts/{id}/withdraw
	
POST	http://localhost:8001/accounts/{id}/deposit
			
POST	http://localhost:8001/accounts



withdraw и deposit имеют body параметр, пример параметра:
{
"amount" : 100 
}