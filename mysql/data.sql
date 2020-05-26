Use pismo_db;

Create table `Account` (id int NOT NULL AUTO_INCREMENT, document_number varchar(200) NOT NULL, available_credit_limit DOUBLE(40, 2) NOT NULL, PRIMARY KEY (id));
Create table `OperationType` (id int NOT NULL AUTO_INCREMENT, description varchar(200) NOT NULL, isNegativeOperator boolean NOT NULL, PRIMARY KEY (id));
Create table `Transaction` (id int NOT NULL AUTO_INCREMENT, account_id int NOT NULL, operation_type_id int NOT NULL, amount DOUBLE(40,2) NOT NULL, event_date timestamp NOT NULL, PRIMARY KEY (id));

delete from OperationType;
Insert into OperationType (description, isNegativeOperator) values ('COMPRA A VISTA', 1);
Insert into OperationType (description, isNegativeOperator) values ('COMPRA PARCELADA', 1);
Insert into OperationType (description, isNegativeOperator) values ('SAQUE', 1);
Insert into OperationType (description, isNegativeOperator) values ('PAGAMENTO', 0);