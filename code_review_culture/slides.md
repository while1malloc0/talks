# Creating a Code Review Culture

<!--
Lead into "Code Review is Useful"
-->

---

# Code Review is useful

---

# Code Review provides
#### code quality improvements

---

# Code Review provides
#### a communication platform

---

# Code Review provides
#### an opportunity to teach

<!--
Teaching this way has immediate practical application
Lead into code review culture is useful
-->

---

# A Code Review *Culture* is useful

---

# Culture

"the set of shared **attitudes**, **values**, **goals**, and **practices** that characterizes an institution or organization" - Mirriam Webster

---

# Agenda:

Examine the practices that contribute to a strong code review culture from the perspective of...

* Organizations
* Authors
* Reviewers

---

# Organizations

<!--
Mark: 3m
-->

---

# Be intentional about your culture

---

## Be intentional about your culture
#### *by communicating the culture*

<!--
Publish your values, goals, and expected attitues
-->

---

## Be intentional about your culture
#### *by establishing a community of experts*

<!--
Language experts should be able to provide advice and feedback about whether a solution is idiomatic both to the language as a whole, and to the best practices of the organization in the use of that language
Domain experts as well as language experts
Domain experts should be able to provide advice and feedback about the correctness of a solution in a particular problem domain
Have a means of becoming a "certified" expert as part of IC career growth
-->

---

## Be intentional about your culture
#### *by developing new experts*

<!--
Experts become gatekeepers if there is no means of making more experts
-->

---

## Be intentional about your culture
#### *by training code reviewers*

<!--
Talk about the fact that code review is a skill that needs to be trained and honed like any other
Lead into code authors
-->

---

# Code Authors

<!--
mark 7m
-->

---

# Make the reviewer's life easier

<!--
Lead into communicating context
-->

---

## Make the reviewer's life easier
#### *by communicating context*

<!--
Recognize that the reviewer probably doesn't have the full context of the problem when reviewing
Lead into images of changelog with no comment
-->

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/no-changelog.jpeg)

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/with-changelog.jpeg)

<!--
Lead into communicating design decisions
-->

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/context-comment.jpeg)

---

## Make the reviewer's life easier
#### *by making the PR the right size*

<!--
Lots of wrong size PRs, very few right size
Obvious thing to get out the way: large PRs are normally bad
Exception: Large deletion PRs
-->

---

## Make the PR the right size
#### *by working in vertical slices*

---

## Vertical slices
#### *ship end-to-end functionality*

<!--
Example: features in a web app
-->

---

## Vertical slices
#### *introduce no new code that is not consumed in the same PR*

---

## Vertical slices
#### *force paring down to an MVP*

---

## Make the reviewer's life easier
#### *by automating the nits*

<!--
Explain what nits are
Automate them using linters, tests, etc
-->

---

## Make the reviewer's life easier
#### *by knowing when to take it offline*

<!--
Explain "take it offline"
If the reviewer leaves a ton of commentary, ask them to pair on addressing comments and making improvements
If the reviewer seems antagonistic, address it offline.
Two types here: misunderstanding and toxicity
-->

---

# Code Reviewers

<!--
mark 17m
-->

---

## Communicate Mutual Respect

<!--
Assume that people can do their jobs
Understand that code review is a pedagogical practice
-->

---

## Communicate Mutual Respect
#### *by knowing when to take it offline*

<!--
If the solution is obviously wrong or would lead you to leave a bunch of comments, set up a pairing session, make a list of revisions to make to the code, and use that time to teach the author
-->

---

## Communicate Mutual Respect
#### *by knowing when you're not the best person to give feedback*

---

## Communicate Mutual Respect
#### *by including justification for critique*

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/comment-no-justification.jpeg)

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/comment-with-justification.jpeg)

---

## Communicate Mutual Respect
#### *by showing that you know that the author can do their job*

<!--
Lead into "question trick"
-->

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/comment-no-question.jpeg)

---

![](https://github.com/while1malloc0/talks/raw/master/code_review_culture/images/comment-with-question.jpeg)

---
## Communicate Mutual Respect
#### *by being as thorough as the PR needs*

<!--
Lead into reviewing in passes
-->

---

# Reviewing in passes

<!--
Mark 23m
-->

---

## Reviewing in passes

#### Each pass is a theme, and some questions to help focus on that theme

---

## Reviewing in passes

#### Make your own. Make a checklist.

---

## Passes to complete every time

#### If there are red flags on any of these, resolve before adding more commentary.

---

## Sizing up

#### What is the general shape of the PR? Is it a completely new feature? Is it a refactor? Is it a one-line change?

---

## Sizing up

#### Is the PR the right size?

<!--
Don't be too picky about the size
-->

---

## Context

#### What is this PR trying to accomplish?

---

## Context
#### Why is this PR trying to accomplish that?

---

## Relevance
#### Is the change made in this PR necessary?

---

## Relevance
#### Does this PR duplicate existing functionality?

---

## Relevance
#### Are there others that should be aware of this PR?

---

## Passes for more in-depth review

#### Do these for more substantial PRs. Pick the ones relevant to the change.

---

## Readability

#### Is the change reasonably understandable by other humans with little/no prior experience of the code?

<!--
Trick: Read the code out loud. If you find yourself saying "why?" maybe something isn't readable
Esoteric language features
-->

---

## Production readiness

#### How do we know when this breaks?

---

## Production readiness

#### Is there new documentation required by this change?


---

## Production readiness

#### Are there tests that prevent regression?

---

## Production readiness

#### Is this change secure?

---

## Naming

#### Do names communicate what things do?

<!--
do they communicate EVERYTHING that they do?
-->

---

## Naming

#### Are the names of things idiomatic to the language?

---

## Naming

#### Do names leak implementation details?

<!--
Don't include mechanical components in names
-->

---

![](https://github.com/while1malloc0/talks/blob/master/code_review_culture/images/bad-variable-name.jpeg)

---

## Gotchas

#### What are the ways in which the added or changed code can break?

---

## Gotchas

#### Is this code subject to any common programming gotchas?

<!--
Common gotchas: transposition errors, off by one errors, memory leaks, null dereferences
-->

---

## Gotchas

#### Is spelling correct and consistent?

<!--
Amazon S3 Java library
-->

---

## Language specific

#### Is the code idiomatic to the language?

---

## Language specific

#### Are new patterns introduced?

---

## Language specific

#### Does the code fall into common pitfalls for the language?

---

## Recap

#### Organizations: be intentional about your code review culture

---

## Recap

#### Authors: make your reviewers' lives easier

---

## Recap

#### Reviewers: be thorough and engage with the pedagogical aspects of code review

<!--
being able to teach is an opportunity. Don't waste it.
-->

---

# Thank you
