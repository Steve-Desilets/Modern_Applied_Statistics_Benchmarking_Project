# Load required libraries
library(readr)
library(pryr)


# Function to perform bootstrapping
bootstrap <- function(data, nBootstraps) {
  bootstrappedMeans <- numeric(nBootstraps)
  n <- length(data)
  set.seed(Sys.time()) # Set seed for reproducibility
  
  for (i in 1:nBootstraps) {
    resample <- sample(data, n, replace=TRUE)
    bootstrappedMeans[i] <- mean(resample)
  }
  
  return(bootstrappedMeans)
}

# Function to calculate confidence interval
confidenceInterval <- function(data, alpha) {
  lower <- quantile(data, alpha / 2)
  upper <- quantile(data, 1 - alpha / 2)
  return(c(lower, upper))
}

# Function to log messages
logMessage <- function(message) {
  cat(message, "\n")
  # Write the log message to a log file
  log_file <- file("C:/Users/steve/modern_applied_statistics_benchmarking_project/modernAppliedStatisticsOutputLogR.txt", open="a")
  writeLines(paste(Sys.time(), message), log_file)
  close(log_file)
}

# Function to test bootstrap
testBootstrap <- function() {
  data <- c(1.0, 2.0, 3.0, 4.0, 5.0)
  nBootstraps <- 1000
  bootstrappedMeans <- bootstrap(data, nBootstraps)
  
  if (length(bootstrappedMeans) != nBootstraps) {
    stop("Test failed: Expected ", nBootstraps, " bootstrapped means, got ", length(bootstrappedMeans))
  } else {
    logMessage("Test passed: Bootstrap function works correctly.")
  }
}

# Load data
data <- read_csv("C:\\Users\\steve\\OneDrive\\Desktop\\Github\\Modern_Applied_Statistics_Benchmarking_Project\\insurance.csv")
  
# Extract charges column
charges <- data$charges
  
# Create vector to store runtimes for benchmarking experiment
runtimes_vector <- c()

# Create vector to store memory used throughout benchmarking experiment
memory_vector <- c()

# Execute 100 iterations of the benchmarking experiment
# Append the runtime for each experimental trial to the runtimes_vector vector
for (iteration in 1:100) {
    
  start_time <- Sys.time()
  
  # Record the amount of memory used before the experimental trials
  memory_before <- pryr::mem_used()
    
  # Perform bootstrapping
  nBootstraps <- 1000
  bootstrappedMeans <- bootstrap(charges, nBootstraps)
    
  # Calculate confidence interval
  alpha <- 0.05
  ci <- confidenceInterval(bootstrappedMeans, alpha)
  lower <- ci[1]
  upper <- ci[2]
  logMessage(paste("95% Confidence Interval for charges:", lower, "-", upper, "\n"))
  
  # Add memory usage to memory_vector
  memory_after <- pryr::mem_used()
  memory_usage <- memory_after - memory_before
  memory_vector <- c(memory_vector, as.numeric(memory_usage))
  
  # Add runtime to runtimes_vector
  end_time <- Sys.time()
  trial_run_time <- end_time - start_time
  runtimes_vector <- c(runtimes_vector, as.numeric(trial_run_time))
  
}

# Let's log the memory usage from the experimental trials
logMessage(paste("Memory usage across all experimental trials:", sum(memory_vector), "bytes"))

# Let's print summary statistics for the trial runtimes gathered
logMessage(paste("Summary of Experimental Trial Runtimes Distribution (in seconds)"))
logMessage(paste("Minimum: ", summary(runtimes_vector)["Min."], " seconds"))
logMessage(paste("1st Quartile: ", summary(runtimes_vector)["1st Qu."], " seconds"))
logMessage(paste("Median: ", summary(runtimes_vector)["Median"], " seconds"))
logMessage(paste("Mean: ", summary(runtimes_vector)["Mean"], " seconds"))
logMessage(paste("3rd Quartile: ", summary(runtimes_vector)["3rd Qu."], " seconds"))
logMessage(paste("Maximum: ", summary(runtimes_vector)["Max."], " seconds"))

# Run the bootstrap test
logMessage(paste("Outcome of the test of the bootstrapping function"))
testBootstrap()


