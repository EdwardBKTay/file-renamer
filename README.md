# File Renamer

File Renamer is a command-line tool written in Go that renames files based on the mappings provided in a CSV file.

## Usage

### Prerequisites

Before using File Renamer, ensure that you have Go installed on your system.

### Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/file-renamer.git
```

Navigate to the project directory:

```bash
cd file-renamer
```

Build the executable binary:

```bash
go build -o file-renamer
```

## Running the Program

File Renamer expects two command-line arguments:

1. `-file`: Path to the CSV file containing old and new file names.
2. `-folder`: Path to the folder containing the files to be renamed.

### Example Usage

```bash
./file-renamer -file=file_mapping.csv -folder=files_to_rename
```

### CSV File Format

The CSV file should contain two columns: the old file names and the new file names. Each row corresponds to a file to be renamed.

Example CSV File:

```csv
old_file_name_1.txt,new_file_name_1.txt
old_file_name_2.txt,new_file_name_2.txt
old_file_name_3.txt,new_file_name_3.txt
```

## Testing

To run tests for the program, execute:

```bash
go test
```
