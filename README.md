# gokriyatest
this project is project for coding challenge kriya - Create simple Api with Go & PostgreeSQL

## How To Run
- First, Clone this project.
- Go to cloned project directory

  example :
    ```bash
       cd /usr/share/gokriyatest
    ```
- Input the postgree configuration in the .env file
  
  example :

    ```bash
       HOST_DATABASE = 192.168.0.100
       PORT_DATABASE = 5432
       USER_DATABASE = postgres
       PASSWORD_DATABASE = my_password12
       NAME_DATABASE = test
    ```
- Run Program
	
    ```bash
       go run .
    ```

## API Collection
- Collection API in Postman : [https://www.getpostman.com/collections/f834c72e863afd8558ab](https://www.getpostman.com/collections/f834c72e863afd8558ab)
- Variabel for Collection :

  example :
    ```bash
       {{server}} => http://192.168.0.151:1234/
       {{token}} => eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDYzMjg0MzMsImlzcyI6IlNpbXBsZSBBcGkgd2l0aCBHbyBcdTAwMjYgUGciLCJVc2VyIjp7ImlkIjoiZjk0MzVkODItMjYyNS0xMWViLWFkYzEtMDI0MmFjMTIwMDAyIiwiZGF0YSI6eyJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBkZW1vLmNvbSIsInBhc3N3b3JkIjoiJDJ5JDEyJDVWM3hqNGlvdmdYdE5EWFlVMkRYWGVTUzFJT25hQUExQ1ZoODNic3hoMUZTQzNkeGtaczhDICIsIlN0YXR1cyI6eyJpc19hY3RpdmUiOnRydWV9fSwicm9sZV9pZCI6IiIsIlJvbGVzIjp7ImlkIjoiMzgxYjc3MDAtZmQyMy00NGI3LTlkMWYtYmVmYmE5ZmE3ZDZhIiwiZGF0YSI6eyJyb2xlX25hbWUiOiJBZG1pbiIsImRlc2NyaXB0aW9uIjoiQWRtaW5pc3RyYXRvciJ9fX19.s8Eo7dAU14dsPdjvZh06L-fZRE7xtaEZctC5veK3kx8
    ```
