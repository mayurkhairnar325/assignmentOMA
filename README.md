# assignmentOMA



# I have created the restaurant ordering application where user can sign and login into an ordering system.
# Once user log’s in, user can use jwt to access the ordering system like user can place or create an order, retrieve order with the help of order id. User can retrieve all the orders as well as user can modify or delete the respective order.

1.	The whole idea behind the order application is First, user will sign-up or sing in with their name, password, email , avatar and  phone number.
curl --location --request POST 'localhost:8000/users/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "mayur",
    "last_name": "k",
    "password": "mayur",
    "email": "m25@icloud.com",
    "avatar": "aaaaaa",
    "phone": "954590"
}
'

2.	Once the user successfully singed in. Then user can login into an ordering application with help of the email and password.
curl --location --request POST 'localhost:8000/users/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"m25@icloud.com",
    "password":"mayur"
}'

Once user logs in us can see this response
{
    "message": "logged in",
    "status": true,
    "user": {
        "ID": 19,
        "CreatedAt": "2022-12-11T21:16:42+05:30",
        "UpdatedAt": "2022-12-11T21:16:42+05:30",
        "DeletedAt": null,
        "first_name": "mayur",
        "last_name": "k",
        "password": "$2a$14$Hj.NpPV4QaQzn8ja3ssDGeU68TtUwvjihyvLlBtGTDZriWIIxcBXi",
        "Email": "m25@icloud.com",
        "avatar": "aaaaaa",
        "phone": "954590",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im0yNUBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjI0YjJkNTNhMTczYjRkNGZhN2NjYzhjMDUzMDZkYTU0IiwiZXhwIjoxNjcwODYwMDAxfQ.2aei7TZFR79pO0W1xAkrU_zYNIm36trJoXPItBhzLlU",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJleHAiOjE2NzEzNzg0MDF9.CzC5MZ-AfxbWkr0_fyVhJfhzPn1qtteUqgFd_A4NR40",
        "user_id": "24b2d53a173b4d4fa7ccc8c05306da54"
    }
}

Using bcrypt to hash the password and its generating two tokens, access token and refresh token
So here I am generating user id with the help of the uuid/ uid

3.	Also we can able to see all the logged in users
curl --location --request GET 'localhost:8000/users' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im00NUBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjA5ZGNlZDE1Mzk1ZDQ5ODFhNDk1N2Y2ZGIwMDc3MWZhIiwiZXhwIjoxNjcxMDI5MzM1fQ.fusTLXiNezaqdXPHwzSi1IUvUDgAmKQPvISeR532-PA'


4.	We retrieve the data of the user with respect to id
curl --location --request GET 'localhost:8000/user/userid/19' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im00NUBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjA5ZGNlZDE1Mzk1ZDQ5ODFhNDk1N2Y2ZGIwMDc3MWZhIiwiZXhwIjoxNjcxMDI5MzM1fQ.fusTLXiNezaqdXPHwzSi1IUvUDgAmKQPvISeR532-PA' \
--data-raw ''


5.	Also we can retrieve respective user data with their user id 
curl --location --request GET 'localhost:8000/user/use/24b2d53a173b4d4fa7ccc8c05306da54' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im00NUBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjA5ZGNlZDE1Mzk1ZDQ5ODFhNDk1N2Y2ZGIwMDc3MWZhIiwiZXhwIjoxNjcxMDI5MzM1fQ.fusTLXiNezaqdXPHwzSi1IUvUDgAmKQPvISeR532-PA'


# To place an order --------

1.	To create a order user has to add the body or data (i am using postman for the same), here it’s mandatory for user to add his user-id while making the request. So we can retrieve that user infromation with help of that user id 
I am using UUID to generate the order ID. 
UUID basically get used to generate random unique Id’s

curl --location --request POST 'localhost:8000/order/orders?token' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im00NUBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjA5ZGNlZDE1Mzk1ZDQ5ODFhNDk1N2Y2ZGIwMDc3MWZhIiwiZXhwIjoxNjcxMDI5MzM1fQ.fusTLXiNezaqdXPHwzSi1IUvUDgAmKQPvISeR532-PA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "pune",
    "menu": "Veg panner, dal makhani",
    "total_items": 2,
    "payment": "Done",
    "user_id":"8f03875f50fa42f68b85b83f40b04eba"
    
}'


Response:
{
    "ID": 13,
    "CreatedAt": "2022-12-13T20:24:45.0848091+05:30",
    "UpdatedAt": "2022-12-13T20:24:45.0848091+05:30",
    "DeletedAt": null,
    "order_id": "488e54838e9c4e80b172e54057f77079",
    "address": "pune",
    "menu": "Veg panner, dal makhani",
    "total_items": 2,
    "payment": "Done",
    "user_id": "8f03875f50fa42f68b85b83f40b04eba"
}


2.	Adding more to this we can retrieve all the order

curl --location --request GET 'localhost:8000/orders' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im0xOEBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjVkOWQ3OWE3YTgzYTQ4Zjk4OTk1ZDlhNDUwMzUzN2EzIiwiZXhwIjoxNjcxMDQwMDU0fQ.IXKrHdiUjSN1Kje6O5zuPxsRUbGbrh4_mYxbsMVaLGQ'



3.	 Retrieve order by ordered id

curl --location --request GET 'localhost:8000/orders/002e55ade33b4e1c8b453b7ecfecb408' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im0xOEBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjVkOWQ3OWE3YTgzYTQ4Zjk4OTk1ZDlhNDUwMzUzN2EzIiwiZXhwIjoxNjcxMDQwMDU0fQ.IXKrHdiUjSN1Kje6O5zuPxsRUbGbrh4_mYxbsMVaLGQ'
, 


4.	we can delete the already existing order, 

curl --location --request DELETE 'localhost:8000/orders/3' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im0xOEBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjVkOWQ3OWE3YTgzYTQ4Zjk4OTk1ZDlhNDUwMzUzN2EzIiwiZXhwIjoxNjcxMDQwMDU0fQ.IXKrHdiUjSN1Kje6O5zuPxsRUbGbrh4_mYxbsMVaLGQ'


5.	And we can modify the already created order. Here i am modifying total items from 2 to 3.

curl --location --request PUT 'localhost:8000/orders/:id' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6Im0xOEBpY2xvdWQuY29tIiwiRmlyc3RfbmFtZSI6Im1heXVyIiwiTGFzdF9uYW1lIjoiayIsIlVpZCI6IjVkOWQ3OWE3YTgzYTQ4Zjk4OTk1ZDlhNDUwMzUzN2EzIiwiZXhwIjoxNjcxMDQwMDU0fQ.IXKrHdiUjSN1Kje6O5zuPxsRUbGbrh4_mYxbsMVaLGQ' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "pune",
    "menu": "Veg panner, dal makhani",
    "total_items": 3,
    "payment": "Done",
    "user_id":"8f03875f50fa42f68b85b83f40b04eba"
    
}'

#In the future scope:
1.	I am planning to fetch the user id direct while creating order. so User don’t need to add that manually 
2.	I will try to fetch the user order with names and phone number,  so user can get to know how much order he/she has been placed from that respective restaurant
3.	I am planning about to add the menu card,  so that user can place the order with available or mentioned menu. If the order is not mentioned in that menu card then user cant able to proceed with the same
4.	I am planning about cart as well, so user can add his/her order into a cart and later on can place the order
5.	Thinking about, invoice controller as well .so one order get placed user can get its invoice.
