# Personal Portfolio

Welcome to my personal portfolio! This website showcases my skills and projects in software development. It's built with a modern tech stack for performance and maintainability.

**Key Features**

* **Dynamic Frontend:** Interactive project displays and engaging presentation of skills using Vue.js.
* **Robust Backend:** Efficient and well-structured backend powered by GoLang (Echo framework).
* **Database Integration:** Seamless data management with PostgreSQL (Sqlc for type-safe queries).
* **Admin Panel:** Streamlined content management for easy portfolio updates.
* **Dockerized:** Easy deployment and portability across environments.
* **Secure Communication:** Website protected with an SSL certificate (e.g., Let's Encrypt) for encrypted data transfer.
* **Automated Renewal:**  Certificate renewal process in place to maintain ongoing website security.
* **NGINX Integration:** Optimized performance with NGINX for reverse proxying and load balancing.
* **CORS Setup:**  Enabled cross-origin resource sharing for seamless frontend-backend interactions.

**Live Demo**

Experience the portfolio in action at [https://maksym-vorobyov.me](https://maksym-vorobyov.me)

**Getting Started**

1. **Prerequisites:**
    * Go (version 1.22.1 or higher)
    * Vue.js (vue/cli 5.0.8 or higher)
    * Node.js (v21.7.2 or higher)
    * Docker
2. **Clone the repository:**
    ```bash
    git clone https://github.com/entl/boutme
    ```
3. **Set up environment variables (if applicable):** 
    * Create a `.env` file to store database credentials, etc.
4. **Build and run with Docker:** 
    ```bash
    docker-compose up -d --build 
    ```

**Project Structure (Brief Overview)**

* **`frontend/`:** Vue.js source code
* **`backend/`:** GoLang source code
* **`docker-compose.yml`:** Configuration for running the app in Docker containers

* [LinkedIn](https://www.linkedin.com/in/maksym-vorobyov/)

**Feedback**

I welcome your feedback and contributions! Feel free to open issues or submit pull requests

**License**
MIT
