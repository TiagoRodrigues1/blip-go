# How the algorithm works

- Collects all the data from the `commits.csv` file.
- Parse the data into a map of string and a structure of a Repo representation
- Calculates the activity score with the given weight for each metric:

| Commits | Files | Lines |
| ------- | ----- | ----- |
|   1.0   | 0.5   | 0.01  |

- Using these values seems like the most balanced approach
    - Commits at 1.0 makes it mostly commit focused, and since the commits will have lower number of occurrences than files and lines we give it a bigger weight.

    - Files are weighted at 0.5, because they typically fall between commits and lines in terms of frequency. We assign it a medium weight to reflect its relative importance.

    - Lines at 0.01, since lines changed is going to have the most occurrences, we give it a smaller weight so that it won't impact the score as much.


# Results

By my definition the most active repositories are:
| Rank | Repository | Activity Score |
| ---- |------------|--------------- |
|1.    | repo250    | 3990.50        |
|2.    | repo518    | 3477.00        |
|3.    | repo476    | 3328.00        |
|4.    | repo127    | 2420.50        |
|5.    | repo126    | 2066.00        |
|6.    | repo795    | 1560.50        |
|7.    | repo742    | 1485.50        |
|8.    | repo381    | 1485.00        |
|9.    | repo740    | 1211.50        |
|10.   | repo259    | 992.50         |

## Pre-requisites

Need GO installed

# How to run
`go run blip.go`