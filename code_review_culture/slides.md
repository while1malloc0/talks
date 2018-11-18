# Creating a Code Review Culture

<!---
- Introduce yourself and say title of talk
- Lead into the uncontroversial statement about code review
-->

---

# Code Review is useful

---

# Code Review provides

* A means of ensuring that code quality remains high
* A place to pass on best practices of an organization
* An easy way for developing engineers to provide technical leadership

<!---
TODO: Break this up into a few slides
TODO: Figure out the transition here. Maybe something about how we maximize the benefits? -->

---

# A Code Review *Culture* is useful

---

# Culture

"the set of shared **attitudes**, **values**, **goals**, and **practices** that characterizes an institution or organization" - Miriam Webster (.com)

---

# Code Review Culture

The set of shared **attitudes**, **values**, **goals**, and **practices** that characterize **an engineering organization's approach to conducting code review.**

<!--- lead in: let's break that down a little bit -->

---

# Attitudes

A shared way of thinking about code review, its importance, and how it should be conducted

---

# Values

A shared understanding of the relative importance of attributes in an implementation

<!---
TODO: Maybe refine this?

- Sometimes best stated as "Given situation x, we use y because z"

-->

---

# Goals

Consensus on what code review is meant to accomplish

<!--- Important point: disagreements on the goal of code review can lead to issues with too many comments -->

---

# Practices

A common set of actions performed by code authors and reviewers as part of a review

<!---
Practices encode the attitudes, goals, and values.
-->

---

# Agenda:

Examine how a code review culture can be fostered by...

* Organizations
* Reviewers
* Authors

---

# Organizations

<!---

-->

---

# Be intentional about your culture

<!---
- Publicize the attitudes, goals, values, and practices
- Communicate about the pedagogical aspects
- Get buy in from leadership
-->

---

# Establish a community of experts

<!---
- Communities should be intentional (?)
- Language experts should be able to provide advice and feedback about whether a solution is idiomatic both to the language as a whole, and to the best practices of the organization in the use of that language
- Domain experts as well as language experts
- Domain experts should be able to provide advice and feedback about the correctness of a solution in a particular problem domain
- Have a means of becoming a "certified" expert as part of IC career growth
-->

---

# Train code reviewers

<!---
- Talk about the fact that code review is a skill that needs to be trained and honed like any other
- Lead into code authors
-->

---

# Code Authors

---

# Make the reviewer's life easier

<!---
- Recognize that the reviewer probably doesn't have the full context of the problem when reviewing
- Key here is communicating with the reviewer. Context is everything
- TODO: This can be expanded into a few slides.
- Communication slides:
- Communicate the context for the change in the PR message (TODO: example of bad and good PR message)
- Highlight areas where design choices were made and show that you've thought about that choice
-->

---
## Make the reviewer's life easier
#### *by communicating context*

---

![](/images/images/no-changelog.png)

---

![](/images/images/with-changelog.png)

---

![](/images/images/context-comment.png)

---
## Make the reviewer's life easier
#### *by making the PR the right size*

---

## Make the PR the right size
#### *by working in vertical slices*

---

## Vertical slices...
#### Ship a complete, end-to-end implementation

---

## Vertical slices...
#### Introduce no new code that is not consumed in the same PR

---

## Make the reviewer's life easier
#### *by automating the nits*

<!---
- Tests, linters, etc. should be run to keep nits to a minimum
-->

---
## Automate the nits
#### *using linters, tests, etc*
---

# Know when to take if offline

<!---
- If the reviewer leaves a ton of commentary, ask them to pair on addressing comments and making improvements
- If the reviewer seems antagonistic, address it offline.
- Two types here: misunderstanding and toxicity
-->

---

# Code Reviewers

---

## Be Thorough and Respectful

<!---
- Assume that people can do their jobs
- Understand that code review is a pedagogical practice
-->

---
## Be Thorough and Respectful
#### *by knowing when to take it offline*

<!---
- If the solution is obviously wrong or would lead you to leave a bunch of comments, set up a pairing session, make a list of revisions to make to the code, and use that time to teach the author
-->

---

## Be Thorough and Respectful
#### *by knowing when you're not the best person to give feedback*

---

## Delivering Feedback
* Give justification for critique that might not have an obvious underlying reason <!--- values -->
* Be straightforward <!--- Avoid "weasel words" -->

<!---
TODO: Break question trick into its own slide
-->
---
## Delivering Feedback

Assume that the author is smart and knows how to do their job.
---

![](/images/images/comment-no-question.png)

---

![](/images/images/comment-with-question.png)

---

# Be as thorough as the PR needs

---

# Reviewing in passes

---

# Passes

Each pass is a theme, and some questions to help focus on that theme

---

# Passes to complete every time

If there are red flags on any of these, resolve before adding more commentary.

---

# Sizing up

* What is the general shape of the PR? Is it a completely new feature? Is it a refactor? Is it a one-line change?

* Is the PR the right size? <!--- don't be nitpicky about this -->

---

# Context

* What is this PR trying to accomplish?
* Why is this PR trying to accomplish that?
* Does this PR actually accomplish what it says it will?
<!---
Hopefully most of this is provided by PR comments
-->

---

# Relevance

* Is the change made in this PR necessary?
* Does this PR duplicate existing functionality?
* Are there others that should be aware of this PR?

---

# Passes for more in-depth review

Do these for more substantial PRs. Pick the ones relevant to the change.

---

# Readability

* Is the change reasonably understandable by other humans with little/no prior experience of the code?

<!---
TODO: Add example slide
Trick: Read the code out loud. If you find yourself saying "why?" maybe something isn't readable
-->

---

# Production readiness

* How do we know when this breaks?
* Is there new documentation required by this change?
* Are there tests that prevent regression?
* Is this change secure?

---

# Naming

* Are variable, function, method names communicative of what they do?
* Are the names of things idiomatic to the language? <!--- most linters should catch this -->
* Do names encapsulate the problem domain? <!--- name after what it does, not what it is -->

---

![](/images/images/bad-variable-name.png)

---

# Gotchas

* What are the ways in which the added or changed code can break?
* Common gotchas: transposition errors, off by one errors, memory leaks, null dereferences
* Is spelling correct and consistent?

---

# Language specific

* Is the code idiomatic?
* Are new patterns introduced?
* Does the code fall into common pitfalls for the language?

---

# Recap

* Organizations: be intentional about your code review culture
* Authors: make your reviewers' lives easier
* Reviewers: be thorough and engage with the pedagogical aspects of code review

---

# Thank you

