Go Anime Trends
A simple and efficient command-line tool written in Go to discover the top trending anime of the current season. It fetches data from a public API and filters the results to show you only the shows with a score of 7.0 or higher.

✨ Features
Fetches all anime airing in the current season (e.g., Summer 2025).

Filters out any show with a community score below 7.0.

Displays a clean, readable list of titles and their scores.

No API keys required.

📋 Prerequisites
Before you begin, ensure you have Go installed on your system. You can download it from the official website: golang.org.

To verify your installation, open a terminal and run:

go version

🚀 Installation & Usage
Clone the repository:
Open your terminal and clone this project to your local machine.

git clone <your-repository-url>

Navigate to the project directory:

cd <your-project-directory>

Run the application:
Execute the following command. Go will automatically handle downloading the necessary dependencies and run the program.

go run .

📝 Example Output
When you run the tool, you will see a list of anime from the current season that meet the criteria, formatted like this:

[*] Fetching current popular animes...

--- Airing Anime This Season (Score > 7.0) ---
#1: Enen no Shouboutai: San no Shou (Score: 7.92)
#2: Lazarus (Score: 7.32)
#3: Wind Breaker Season 2 (Score: 7.80)
#4: Witch Watch (Score: 7.47)
#5: Tu Bian Yingxiong X (Score: 8.54)
#6: Vigilante: Boku no Hero Academia Illegals (Score: 7.57)
#7: Kowloon Generic Romance (Score: 7.57)
#8: Kijin Gentoushou (Score: 7.52)
#9: Aharen-san wa Hakarenai Season 2 (Score: 7.44)
... and more!

(Note: The list above is a representative example. The actual output will reflect the anime trending in the current season at the time you run the script.)

Happy Watching! 📺