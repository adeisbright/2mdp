# 2 Million Debtors Program (2MDP)

2 Million Debtors is a  Command Line Program with components and features aimed at fostering the understanding of system internals and how to go low with languages. 


In this program, we will generate a database of 200 million customers with an assumed amount they are owing. 
Name and amount will be automatically generated.  The duration for repayment and the date they collected payment will also be automatic

After the generation or during the generation, the names will need to be stored in a file. 

After storing the file, our program will then read from the file, and calculate the date each customer ought to pay back their debt, and then update the record, 

Now, the unique id for each  customer will also be hashed and can only be read during the process of creating  or updating the list. 

This project will cover the following : 
1. Memory Utilisation 
1. Data Structures 
1. Algorithm 
1. I/O Operations
1. Security 
1. Data Encryption 

I also aim to be able to show the difference between making disk calls to read from  a file , and reading directly from cache (in-memory store) 

## System Components 

The identified components we will be building for this program are : 
1. Customers 
2. Data Manager (File Processor)
3. Cache Manager 

The customers component will hold data and how to process data belonging to customers. 

Data Manager helps us to read, write to , and update the file that serves as our persistent data storage 

Cache Manager will help us in interacting with the memory


