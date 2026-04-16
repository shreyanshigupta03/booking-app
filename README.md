
# 🎟️ Go Conference Ticket Booking CLI App

A simple command-line application built using **Go (Golang)** that allows users to book tickets for a conference.
This project demonstrates core Go concepts like structs, slices, concurrency (goroutines), and synchronization using WaitGroups.

---

## 🚀 Features

* Book conference tickets via CLI
* Input validation (name, email, ticket count)
* Tracks remaining tickets
* Stores user booking data
* Sends tickets asynchronously using goroutines
* Uses WaitGroup for concurrency management
* Displays list of booked users

---

## 🛠️ Tech Stack

* Language: Go (Golang)
* Concepts:

  * Structs
  * Slices
  * Functions
  * Goroutines
  * WaitGroups (`sync` package)
  * Modular code (`helper` package)

---

## 📂 Project Structure

booking-app/
│── main.go
│── helper/
│   └── helper.go
│── README.md

---

## ▶️ How to Run

### 1. Clone the repository

git clone [https://github.com/your-username/booking-app.git](https://github.com/your-username/booking-app.git)
cd booking-app

### 2. Run the application

go run .

---

## 🧑‍💻 How It Works

1. User enters:

   * First Name
   * Last Name
   * Email
   * Number of tickets

2. Input is validated using helper functions.

3. If valid:

   * Tickets are booked
   * User data is stored
   * Ticket is sent asynchronously (after delay)

4. App continues until all tickets are sold out.

---

## ⚙️ Example Output

Welcome to Go conference booking application!
We have total 50 tickets and 50 tickets are still remaining.

Enter your first name:
John
Enter your last name:
Doe
Enter your email address:
[john@example.com](mailto:john@example.com)
Enter your number of tickets you want to book:
2

Thank you John Doe for booking 2 tickets.
48 tickets are remaining.

Sending tickets:
2 tickets for John Doe to email address [john@example.com](mailto:john@example.com)

---

## 🧠 Learning Highlights

* Understanding goroutines for async execution
* Managing concurrency using sync.WaitGroup
* Working with structs and slices in Go
* Handling user input in CLI applications

---

## 📌 Future Improvements

* Add persistent storage (database or file)
* Convert to web-based application
* Integrate real email service
* Generate unique ticket IDs
* Add unit tests

---

## 🤝 Contributing

Contributions are welcome! Feel free to fork this repo and submit a pull request.

---




