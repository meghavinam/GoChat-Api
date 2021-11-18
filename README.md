### database table creation


CREATE TABLE `chat_messages` ( `id` INT NOT NULL AUTO_INCREMENT , `message` TEXT NOT NULL , `created_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP , `updated_date` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`id`)) ENGINE = InnoDB;


#### Virtualhost

<VirtualHost *:80>
ProxyPreserveHost On
ProxyRequests Off
ServerName checkapi.abc.com
ProxyPass / http://localhost:8086/
ProxyPassReverse / http://localhost:8086/
</VirtualHost>
 

#### Extra installation
add proper configurations in src/config/config.json file
#### run and build the fiile

run 
sh buildfile.sh 

    





	
