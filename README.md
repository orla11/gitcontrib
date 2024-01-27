# GitContrib

GitContrib is a command-line tool written in Go that helps you track and visualize your Git contributions over the last six months. It provides functionality to scan for Git repositories in specified folders and generates a graph showing the number of commits made each day.

## Table of Contents

-   [GitContrib](#gitcontrib)
    -   [Table of Contents](#table-of-contents)
    -   [Installation](#installation)
    -   [Usage](#usage)
        -   [Adding a New Folder for Scanning](#adding-a-new-folder-for-scanning)
        -   [Viewing Git Contributions Stats](#viewing-git-contributions-stats)

## Installation

To use GitContrib, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/GitContrib.git
    ```

2. Change into the project directory:

    ```bash
    cd GitContrib
    ```

3. Build the application:

    ```bash
    go build cmd/main/main.go
    ```

4. Run the application:
    ```bash
    ./main
    ```

## Usage

### Adding a New Folder for Scanning

You can add a new folder for scanning Git repositories using the following command:

```bash
./main -add /path/to/folder
```

This command will recursively scan the specified folder and its subfolders, searching for Git repositories. The discovered repositories will be added to the list of tracked repositories.

### Viewing Git Contributions Stats

To view Git contributions stats, use the following command:

```bash
./main -email your@email.com
```

Replace `your@email.com` with the email address associated with your Git commits. The tool will generate a graph showing the number of commits made each day over the last six months.

This project is created for learning/testing purposes - Original project can be found [here](https://gist.github.com/flaviocopes/bf2f982ee8f2ae3f455b06c7b2b03695).
