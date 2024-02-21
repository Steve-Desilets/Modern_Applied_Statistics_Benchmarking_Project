# Modern Applied Statistics Benchmarking Project

For the project contained within this repository, we put ourselves in the position of a data scientist working for a company whose leadership team is interested in evaluating the financial and operational benefits of switching from R to Golang for data science programs that will be run by cloud providers. For this study, the company's leadership team requests that the data scientist use Golang and R to implement one of the eight most important statistical ideas of the past 50 years (as identified by Dr. Andrew Gellman and Dr. Aki Vehtari), which are: 

1) Counterfactual causal inference
2)  Bootstrapping and simulation-based inference
3)  Overparameterized models and regularization
4)  Bayesian multilevel models
5)  Generic computation algorithms
6)  Adaptive decision analysis
7)  Robust inference
8)  Exploratory data analysis

The company's leadership team requests that the data scientist conduct benchmarking experiments to understand the memory usage and computational speed with which Golang and R can complete identical data science tasks. Based on these measurements, leadership requests that the data scientist estimate the financial implications of selecting Golang or R for such data science computing operations. In addition, the company's leadership team also wants to understand whether both languages can be leveraged for logging, testing, and profiling for these types of experiments. The company's leadership team requests that the data scientist deliver a report outlining findings and recommendations related to all of these subjects.
   






For the project contained within this repository, we put ourselves in the position of a data scientist working for a company whose leadership team is interested in evaluating whether it is over-paying the data scientists, data engineers, and data analysts at the company.  Specifically, the company's leadership points out that recent advancements in natural language processing (NLP) have led to the release of many new tools that can write or edit entire programs for data scientists.  Given that such tools have been streaming onto the market, the company's leadership wonders whether the company still truly needs to be paying $120,000 and $175,000 annual salaries to junior Golang programmers and senior software engineers, respectively.  The company's leadership requests that the data scientist learn more about three newly developed automated programming tools: 1) an Automated Code Generator , 2) an AI-Assisted Programming Tool, and 3) a large language model (LLM) capable of generating code (AI-Generated code). The data scientist must report to the company's leadership about the effectiveness of each of these tools and the degree to which the company may be able to save money on programmer salaries due to these tools automating the work associated with these positions.

For this evaluation, the data scientist focuses on leveraging these three automated programming tools to calculate linear regression models for each of the four Anscombe Quartet datasets.  For the first (Automated Code Generator) phase of the experiment, the data scientist examines the extent to which he's able to leverage the Golang Generate command to write the regression-calculating program.  For the second (AI-Assisted Programming) phase of the experiment, the data scientist evaluates the extent to which the Github Copilot application is able to improve the performance of a previously written Golang program focused on calculating regression models for Anscombe Quartet data (Desilets 2024).  For the third (AI-Generated Code) phase of the experiment, the data scientist requests that ChatGPT write the entire program for him and saves the original text of the conversation with the LLM in this repository as a separate text file. For each of these three automated programming methods, the data scientist also performs 100 iterations of the regression-calculating experiments and records the runtime for each experimental trial.  The data scientist then prints the total and average runtime for these 100 experimental trials to an output text file - each of which is included within this repository.

After leveraging Golang Generate for the first (Automated Code Generator) evaluation,  the data scientist observed the strengths and weaknesses of this command for reducing the programming workload of data scientists.  The data scientist created the "AutomatedCodeGenerationCode.go" file within this repository from scratch using Go Generate.  Interestingly, as the data scientist wrote the program, Go Generate would predict each line of code that he was trying to write in the same way that Word and Gmail autocomplete features aim to predict users' sentences.  In fact, if the data scientist typed comments explaining what he wanted the subsequent lines of code to complete, then Go Generate could accurately predict entire blocks of code he wanted to write. Furthermore, this command seemed intelligent enough to incorporate outside information into its predictions.  (Specifically, Go Generate generated code defining all the data points in the Anscombe Quartet just because the data scientist indicated in a comment that he wanted to work with Anscombe Quartet data). Despite these successes, Go Generate did still have some slight weaknesses.  Specifically, Go Generate's predictions were not always fully accurate and Go Generate still did rely on the data scientist to oversee the overall structure of the code quite a bit as the file was being developed.  In the end, this program, correctly calculated the regression lines as y = 0.5*x + 3, and completed these calculations with an average experimental trial runtime of 2,073 microseconds - the slowest runtime key performance indicator (KPI) resulting from our three evaluations.  Though the Go Generate command may not have yielded the strongest results computationally or have written the program entirely independently, this program was probably the second most successful methodology in reducing the overall programming workload for the data scientist.

The data scientist then leveraged the Github Copilot application in the Goland Integrated Development Environment for the evaluation of an AI-Assisted Programming Tool.  Specifically, the data scientist took a previously written program designed to complete the same Anscome Quartet regression line calculation task and utilized Github Copilot to review and improve the code (Desilets 2024). The resulting code is included in this repository in a file named "AI_assisted_code.go."  Github Copilot for Goland currently functions in a similar way to the Editor Tool in Microsoft Word.  The data scientist clicked on every line of code and asked Github Copilot whether it had any recommendations for how to improve it.  Github Copilot did provide a few helpful suggestions. For example, Github Copilot suggested a couple annotations that improved the code's readability and documentation.  Github Copilot also suggested adding a few lines of code to the "LinearRegression" function that would return an error if the function encountered an empty data set - preventing the code from panicking in such a situation.  While Github Copilot seems like it may have more advanced features available on the horizon in Goland, those features are still prototypes that are not yet available to the entire public. Since this application largely just acts as an editor in Goland at the moment, it was probably the least effective of at reducing the programming workload of the three methodologies evaluated. However, the program still correctly calculated each of the regression lines with an average experimental trial runtime of 503 microseconds - the second fastest runtime KPI observed.

The data scientist then requested that OpenAI's ChatGPT generate code for the third evaluation - the evaluation of AI-Generated code.  The data scientist included the entire text of this conversation with ChatGPT in this repository's "Full_Text_of_Conversation_with_ChatGPT_for_aiGeneratedCode.txt" file and the resulting Goland program in this repository's "aiGeneratedCode.go" file. Writing this entire program in just seconds, ChatGPT was wildly successful as a tool in reducing the workload for the data scientist. Moreover, the Goland program correctly calculated each of the regression lines and the resulting average experimental trial runtime was 25 microseconds.  In this way, ChatGPT's code was also far more computationally efficient than the code generated by the first two methods - which each required more human intervention as well.

If I were to make a recommendation to management based on the findings of this study, I would encourage leadership to consider adjusting its payscale for data scientists given that so many newly available tools can reduce the workload for programmers.  In particular, this evaluation displayed that ChatGPT's ability to write AI-Generated code nearly fully automates the job of a programmer and results in more efficient code than would otherwise be created with more human intervention. Consequently, the company should be able to hire fewer individuals who may be less experienced / educated and pay them lower salaries as a result, given that ChatGPT has the ability to complete much of the position requirements on their behalf.



References

Choi, Miri. 2018. "Medical Cost Personal Datasets". Kaggle. https://www.kaggle.com/datasets/mirichoi0218/insurance

Gelman, Andrew, and Aki Vehtari. 2022. “What Are the Most Important Statistical Ideas of the Past 50 Years?" Journal of the American Statistical Association, 116: 2087–2097. 

OVHcloud. 2024. "High-performance cloud solutions for the best possible prices". OVHcloud. https://us.ovhcloud.com/public-cloud/prices/ 
