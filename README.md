## 📑 About the Project
<p align="justify">Immersive Dashboard App<br>
  <br>
This RESTful API was developed by using Golang and written based on Clean Architecture principles. Built with Echo as web framework, GORM as ORM, MySQL as DBMS etc.
</p>

## 🛠 Tools
**Backend:** <br>
![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)
![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)

**Deployment:** <br>
![Google Cloud](https://www.google.com/url?sa=i&url=https%3A%2F%2Fseeklogo.com%2Fvector-logo%2F336116%2Fgoogle-cloud&psig=AOvVaw3o-wCICivwwkA-Tk90-PoJ&ust=1686311725449000&source=images&cd=vfe&ved=0CBEQjRxqFwoTCKio1OzOs_8CFQAAAAAdAAAAABAE&)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge&logo=ubuntu&logoColor=white)

**Communication:**  
![GitHub](https://img.shields.io/badge/github%20Project-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-%237289DA.svg?style=for-the-badge&logo=discord&logoColor=white)

# 🔗 ERD
<img src="ERD.jpg">

# 🔥 Open API

Simply [click here](https://app.swaggerhub.com/apis/dimasyudhana/immersivedashboardproject/1.0.0) to see the details of endpoints we have agreed with our FE team.

<details>
  <summary>👶 Users </summary>
  
| Method      | Endpoint            | Params      |q-Params            | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- |--------------------| ----------- | --------------------------------------- |
| POST        | /register           | -           |-                   | NO          | Register a new Use                      |
| POST        | /login              | -           |-                   | YES         | Login to the system                     |
| GET         | /users              | -           |-                   | YES         | Show user profile                       |
| PUT         | /users              | -           |-                   | YES         | Update user profile                     |
| DELETE      | /users              | -           |-                   | YES         | Update user profile                     |


  
</details>

<details>
  <summary>📑 Classes</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /classes            | -           | YES         | Post a books                            |
| GET         | /classes            | -           | YES         | Get All book                            |
| GET         | /classes            | -           | YES         | Get MyBooks                             |
| PUT         | /classes            | class_id    | YES         | Edit book                               |
| DELETE      | /classes            | class_id    | YES         | Delete book                             |
| GET         | /classes            | class_id    | YES         | Get books Detail                        |  

  </details>
     <details>
  <summary>📠 Mentees</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /mentees            | -           | YES         | Make User Rent detail                   |
| GET         | /mentees            | mentee_id   | YES         | Get User Rent detail                    |


  </details>
  <details>
   <summary>🔊 Feedbacks</summary>
  
| Method      | Endpoint            | Params      | JWT Token   | Function                                |
| ----------- | ------------------- | ----------- | ----------- | --------------------------------------- |
| POST        | /feedbacks          | -           | YES         | Make User Rent                          |
| GET         | /feedbacks          | rents_id    | YES         | Get User Rent                           |
| GET         | /feedbacks          | -           | YES         | Get History Rent from User              |
| GET         | /feedbacks          | -           | YES         | Get History Book Rented from User       |


  </details>
    
 

# 🛠️ How to Run Locally

- Clone it

```
$ git clone https://github.com/GroupProject2-Kelompok4/BE.git
```

- Go to directory

```
$ cd ./GroupProject2-Kelompok4/BE
```
- Run the project
```
$ go run main.go
```

- Voila! 🪄

### 🧰Backend

- [Github Repository for the Backend team](https://github.com/GroupProject2-Kelompok4/BE.git)
- [Swagger OpenAPI](https://app.swaggerhub.com/apis/dimasyudhana/immersivedashboardproject/1.0.0)


# 🤖 Author

-  Dimas A Yudhana  <br>  [![GitHub](https://img.shields.io/badge/Dimas-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)](https://github.com/dimasyudhana)

<h5>
<p align="center">Created by Group 4 ©️ 2023</p>
</h5>