# goassignment

least recently used cache(LRU cache) using Go language

# summary of what this code does

This code creates a basic LRU cache with a specified capacity. When an item is added to the cache (Set method), it’s added to the front of a doubly linked list and a reference to the list element is stored in a map. If the cache is at capacity, the least recently used item (the last item in the list) is removed. When an item is retrieved from the cache (Get method), it’s moved to the front of the list. The Delete method removes an item from the cache.

The Set method also accepts a duration parameter. If this parameter is greater than 0, it creates a timer that will automatically delete the item from the cache after the specified duration.

# INSTRUCTION TO RUN THIS CODE

INSTALL GO vrsion above go1.21.6
RUN COMMAND IN TERMINAL TO DOWNLOAD DEPENDENCIES-> go mod tidy
RUN COMMAND IN TERMINAL TO RUN A PROGRAM(Note:run it from main.go file directory)-> go run main.go

# REACT dependency

# Run below commands in terminal

cd LRUcache-react-app
npm install
npm run dev

# Repo link-> https://github.com/abhilasha336/goassignment

# POSTMAN COLLECTION FILENAME-> LRU Cache Endpoints.postman_collection.json
