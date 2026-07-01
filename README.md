# Nginx & Golang Round-Robin Load Balancer

This project demonstrates how to configure an **Nginx** reverse proxy to act as a load balancer for **three concurrent Golang backend servers**.


![Project Diagram](https://levelup.gitconnected.com/nginx-the-underrated-load-balancer-that-might-already-be-running-in-your-stack-b56136301eb1)
---

## System Architecture

The architecture consists of one entry point (Nginx) listening on port 80 and distributing requests across three Go application instances.

* **Load Balancer:** Nginx (Port 80)
* **Backend Server 1:** Golang App (Port 8081)
* **Backend Server 2:** Golang App (Port 8082)
* **Backend Server 3:** Golang App (Port 8083)

---

## Load Balancing Algorithm: Round-Robin

The project utilizes the default Nginx load balancing algorithm: **Round-Robin**.

### How it Works
* Requests are distributed **sequentially** and **cyclically** down the list of servers.
* **Request 1:** Directed to Server 1 (Port 8081).
* **Request 2:** Directed to Server 2 (Port 8082).
* **Request 3:** Directed to Server 3 (Port 8083).
* **Request 4:** Loops back to Server 1 (Port 8081).

### Key Features
* **Equal Distribution:** Every server handles an identical share of the total traffic.
* **Stateless:** The proxy does not need to calculate server loads or response times.
* **Automatic Failover:** If one Go server crashes, Nginx automatically skips it and routes traffic to the remaining healthy servers.