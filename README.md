# tv-show-tracker
The TV Show Tracker web app will allow users to search for TV shows, view details about their favorite series (e.g., seasons, episodes, ratings), and track their viewing progress. This project will integrate external APIs to fetch show data, support full CRUD operations for user-managed lists, and feature a modern, responsive UI.

## Core Goals:

- Search & Fetch Show Data:

    Integrate with an external TV show API (e.g., TV Maze API) to search and fetch show details.

- User Watchlist Management:

    Allow users to add, update, and delete shows from their personal watchlist.
    Track viewing progress (e.g., mark episodes as watched, track season progress).

- UI/UX & Interactivity:

    Build a clean, responsive UI to display show information, user watchlists, and progress tracking.
    Implement filtering, sorting, and search capabilities on the watchlist.

- Deployment & Documentation:

    Deploy the app using a platform like Railway or Heroku.
    Polish documentation to highlight deployment processes and lessons learned.

## Project Breakdown by Day
### Define Scope & Set Up Project

Goals:

    Plan the app’s features, user flows, and database schema.
    Set up the project folder and initialize the Go module with necessary dependencies.

Tasks:

    Define the feature list:
        External API integration to search for TV shows.
        CRUD operations for user watchlists.
        Tracking viewing progress (watched/unwatched episodes).
        Filtering and sorting the watchlist.
    Sketch a UI wireframe (show search bar, watchlist table, detail view for shows).
    Set up the project folder structure:

    /tv-show-tracker
      ├── /templates        // HTML templates
      ├── /static           // CSS/JS files
      ├── main.go
      ├── routes.go
      ├── models.go
      └── database.go

    Run go mod init tv-show-tracker and install dependencies: Gin, GORM, Resty, godotenv.

📌 Deliverable: Project initialized with clear scope and folder structure.

### Set Up Database & Models

Goals:

    Define models for storing TV shows and user watchlists.
    Set up the database using GORM with either SQLite or PostgreSQL.

Tasks:

    Design models for:
        Show: ID, Title, Summary, Rating, Image URL, etc.
        Watchlist: UserID (if multi-user), ShowID, LastWatchedEpisode, Status (e.g., "Watching", "Completed").
    Write migration functions to create the necessary tables.
    Seed some sample data (if desired) to test display functionality.

📌 Deliverable: A fully structured database schema and working GORM models.

### Implement External API Integration & Backend CRUD

Goals:

    Build API endpoints to search for TV shows using an external API.
    Implement CRUD operations for managing the watchlist.

Tasks:

    Use Resty to fetch data from an external TV show API (e.g., TV Maze API):
        Create an endpoint (e.g., GET /search) that accepts a query parameter and returns matching shows.
    Implement backend endpoints using Gin:
        POST /watchlist to add a show.
        GET /watchlist to retrieve the user’s watchlist.
        PUT /watchlist/:id to update watch progress.
        DELETE /watchlist/:id to remove a show.
    Validate inputs and handle errors appropriately.

📌 Deliverable: Fully functional API endpoints for searching TV shows and managing the watchlist.

### Build the Frontend & Connect to Backend

Goals:

    Create dynamic HTML templates for displaying show search results and the user’s watchlist.
    Connect the frontend with backend endpoints.

Tasks:

    Develop HTML templates for:
        Homepage: Includes a search bar and displays search results.
        Watchlist Page: Lists added shows with options to update progress or remove a show.
        Show Detail Page: Displays detailed information fetched from the external API.
    Use Go’s html/template for rendering dynamic content.
    Implement AJAX or simple form submissions to interact with your API endpoints.
    Use basic CSS (or a framework like TailwindCSS) to style the pages for a clean look.

📌 Deliverable: A functional, user-friendly frontend that communicates with the backend.

### Enhance UI/UX & Add Interactive Features

Goals:

    Improve user experience by adding filtering, sorting, and progress tracking functionalities.
    Refine UI elements based on user feedback or personal preference.

Tasks:

    Add UI components to filter or sort the watchlist (e.g., by rating, status).
    Implement interactive elements, such as a progress bar for each show or a modal popup for show details.
    Update backend logic to support these interactive features.
    Polish CSS styling and ensure responsiveness across devices.

📌 Deliverable: An engaging and interactive TV Show Tracker UI with advanced features.

### Testing, Debugging & Final Touches

Goals:

    Thoroughly test the full stack: API integration, database operations, and UI interactions.
    Debug issues and optimize performance.

Tasks:

    Test API endpoints with tools like Postman.
    Manually test the frontend: search functionality, watchlist updates, filtering, and sorting.
    Fix any bugs encountered and optimize SQL queries if necessary.
    Finalize documentation, including API usage and deployment instructions.

📌 Deliverable: A fully tested, polished app with smooth interactions and error handling.

### Final Testing, Documentation & Deployment

Goals:

    Perform final end-to-end testing.
    Prepare comprehensive documentation.
    Deploy the app using a platform like Heroku or Railway.

Tasks:

    Conduct comprehensive testing of all app features.
    Update your README, including setup, usage instructions, and lessons learned.
    Use godotenv to manage environment variables.
    Deploy the app to Railway, Heroku, or another hosting platform.
    Record a short demo video and capture screenshots for content creation.

📌 Deliverable: The TV Show Tracker is live, documented, and ready to be showcased in your portfolio and on social media.

