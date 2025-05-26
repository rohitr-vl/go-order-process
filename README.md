# go-order-process
Golanag practice - order processing system
### **🧪** **Problem Title:**  **Real-Time Order Processing System**

---
### **Problem Statement:**

You are tasked with building a **Real-Time Order Processing System** for an e-commerce platform backend. The system should simulate order creation, processing, and status updates using **Go**, and should involve the following components:

1. **Order Producers** (simulating new orders via a REST API or scheduled task)
    
2. **A Queue System** (for decoupled processing)
    
3. **A Pub/Sub Mechanism** (for notifying services of order status changes)
    
4. **A Database** (for persisting and querying order data)
    

---

### **✅** **Requirements:**

  #### **1.** **Order API (Producer)**
- Implement a RESTful API endpoint: POST /orders
    
- Accept JSON payload:
```JSON
{
  "customer_id": "string",
  "items": ["item1", "item2"],
  "total": 123.45
}
```
- On receiving an order, **enqueue** it into a message queue (e.g., channel, Redis, NATS, or Kafka – based on choice and simplicity)

#### **2.** **Order Worker (Consumer)**
- A Go routine or service that consumes orders from the queue
    
- For each order:
    
    - Validate the payload
        
    - Store order with status processing in the database
        
    - Simulate processing time (e.g., sleep for 1-3 seconds)
        
    - Update the order status to completed
        
    - **Publish** an event to a pub/sub topic: "order.completed"
#### **3.** **Notification Service (Subscriber)**

- Subscribes to the "order.completed" topic
    
- Logs or sends (mock) a notification like:
```
Email sent to customer <customer_id> for Order #<order_id> - Status: completed
```
#### **4.** **Database**

- Use any relational (e.g., PostgreSQL) or NoSQL (e.g., MongoDB) database
    
- Store fields like:
```
order_id, customer_id, items (JSON/text), total, status, created_at, updated_at
```
- Implement an endpoint: GET /orders to return all stored orders (filterable by status)


### **🧰** **Tech Stack Recommendations (can vary slightly):**

- **Go** (standard library + goroutines/channels)
    
- **Redis/NATS/Kafka** for Queue and Pub/Sub (or even in-memory for simplicity if needed)
    
- **PostgreSQL or SQLite** for storage (using GORM or database/sql)
    
- **Gin** or **Echo** for HTTP server
    
- **Docker** for running services locally

### **🎯** **Goals for Assessment:**

- Understanding of system design with queues and decoupling
    
- Use of concurrency (goroutines, channels)
    
- Integration with pub/sub systems
    
- Clean architecture and modular code
    
- Efficient and safe DB handling (transactions, connection pooling)
    
- Logging and error handling practices