# ShakeSearch

Welcome to the Pulley Shakesearch Take-home Challenge! In this repository,
you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

You can see a live version of the app at
https://pulley-shakesearch.herokuapp.com/. Try searching for "Hamlet" to display
a set of results.

In it's current state, however, the app is just a rough prototype. The search is
case sensitive, the results are difficult to read, and the search is limited to
exact matches.

## Your Mission

Improve the search backend. Think about the problem from the **user's perspective**
and prioritize your changes according to what you think is most useful. 

## Evaluation

We will be primarily evaluating based on how well the search works for users. A search result with a lot of features (i.e. multi-words and mis-spellings handled), but with results that are hard to read would not be a strong submission. 


## Submission

1. Fork this repository and send us a link to your fork after pushing your changes. 
2. Heroku hosting - The project includes a Heroku Procfile and, in its
current state, can be deployed easily on Heroku's free tier.
3. In your submission, share with us what changes you made and how you would prioritize changes if you had more time.

## My instance of Shakesearch
https://searchbar7.herokuapp.com


## My approach
User is interested in getting the queried word, but the context of the word is important as well.
My take is to use regular expressions to bring the full paragraph where the word is at and to bring recommendations to the user.
With word recommendations you can recommend instead of correcting the user input.
3 regular expressions were used, one for normal strings (i.e. "Hamlet"), another regexp for words with just uppercase letters, and the last one to find words that begin with the users input.
If I'd have more time to make a robust backend for searches I would prioritize the robustness of my regular expressions in order to handle non-case-sensitive searches, I'd limit the regular expression to retrieve 2 or 3 paragraphs of the conversation where the word is at, and to make the infra of the backend a little bit better to avoid lots of round trips for the recommendations endpoint. On the front I would implement something like the Levenshtein distance in order to handle misspellings in case the user wants it.
