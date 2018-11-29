# Creating a Code Review Culture

Note:
Lead into "Code Review is Useful"

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

Note:
Teaching this way has immediate practical application
Lead into code review culture is useful

---

# A Code Review *Culture* is useful

---

# Culture

"the set of shared **attitudes**, **values**, **goals**, and **practices** that characterizes an institution or organization" - Mirriam Webster (.com)

---

# Code Review Culture

The set of shared **attitudes**, **values**, **goals**, and **practices** that characterize **an engineering organization's approach to conducting code review.**

---

# Attitudes

A shared **way of thinking** about code review

---

# Values

A shared **ranking** of what's important in code

---

# Goals

A shared **agreement** on the **purpose of code review**

---

# Practices

A shared **set of actions** performed as part of code review

Note:
Practices can encode attitudes, goals, and values
Sneakily influence culture

---

# Agenda:

Examine the practices that contribute to a strong code review culture from the perspective of...

* Organizations
* Authors
* Reviewers

---

# Organizations

Note:
TODO: Figure out better description for "middle-out" change that doesn't use the words "Cross functional working group"
---

# Be intentional about your culture

---
## Be intentional about your culture
#### *by communicating the culture*

Note:
Publish your values, goals, and expected attitues
---
## Be intentional about your culture
#### *by establishing a community of experts*

Note:
- Language experts should be able to provide advice and feedback about whether a solution is idiomatic both to the language as a whole, and to the best practices of the organization in the use of that language
- Domain experts as well as language experts
- Domain experts should be able to provide advice and feedback about the correctness of a solution in a particular problem domain
- Have a means of becoming a "certified" expert as part of IC career growth

---
## Be intentional about your culture
#### *by developing new experts*

Note:
Experts become gatekeepers if there is no means of making more experts
---
## Be intentional about your culture
#### *by training code reviewers*

Note:
- Talk about the fact that code review is a skill that needs to be trained and honed like any other
- Lead into code authors

---

# Code Authors

---

# Make the reviewer's life easier

Note:
- Lead into communicating context

---
## Make the reviewer's life easier
#### *by communicating context*

Note:
- Recognize that the reviewer probably doesn't have the full context of the problem when reviewing
- Lead into images of changelog with no comment

---

![](/images/no-changelog.png)

---

![](/images/with-changelog.png)

Note:
- Lead into communicating design decisions

---

![](/images/context-comment.png)

---
## Make the reviewer's life easier
#### *by making the PR the right size*

Note:
- Lots of wrong size PRs, very few right size
- Obvious thing to get out the way: large PRs are normally bad
- Exception: Large deletion PRs

---

## Make the PR the right size
#### *by working in vertical slices*

---

## Vertical slices
#### *ship end-to-end functionality*

Note:
Example: features in a web app

---

## Vertical slices
#### *introduce no new code that is not consumed in the same PR*

---

## Vertical slices
#### *force paring down to an MVP*

---

## Make the reviewer's life easier
#### *by automating the nits*

Note:
- Explain what nits are
- Automate them using linters, tests, etc

---

## Make the reviewer's life easier
#### *by knowing when to take it offline*

Note:
- Explain "take it offline"
- If the reviewer leaves a ton of commentary, ask them to pair on addressing comments and making improvements
- If the reviewer seems antagonistic, address it offline.
- Two types here: misunderstanding and toxicity

---

# Code Reviewers

---

## Communicate Mutual Respect

Note:
- Assume that people can do their jobs
- Understand that code review is a pedagogical practice

---
## Communicate Mutual Respect
#### *by knowing when to take it offline*

Note:
- If the solution is obviously wrong or would lead you to leave a bunch of comments, set up a pairing session, make a list of revisions to make to the code, and use that time to teach the author


---

## Communicate Mutual Respect
#### *by knowing when you're not the best person to give feedback*

---
## Communicate Mutual Respect
#### *by including justification for critique*

---

![](/images/comment-no-justification.png)

---

![](/images/comment-with-justification.png)

---
## Communicate Mutual Respect
#### *by showing that you know that the author can do their job*

Note:
- Lead into "question trick"

---

![](/images/comment-no-question.png)

---

![](/images/comment-with-question.png)

---
## Communicate Mutual Respect
#### *by being as thorough as the PR needs*

Note:
- Lead into reviewing in passes

---

# Reviewing in passes

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

Note:
- Don't be too picky about the size
---
## Context

#### What is this PR trying to accomplish?
---
## Context
#### Why is this PR trying to accomplish that?
---
## Context
#### Does this PR actually accomplish what it says it will?

Note:
Hopefully most of this is provided by PR comments
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
Note:
- TODO: Add example slide
- Trick: Read the code out loud. If you find yourself saying "why?" maybe something isn't readable

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
Note:
- do they communicate EVERYTHING that they do?
---
## Naming
#### Are the names of things idiomatic to the language?
---
## Naming
#### Do names leak implementation details?

Note: Don't include mechanical components in names
---

![](/images/bad-variable-name.png)

---
## Gotchas

#### What are the ways in which the added or changed code can break?
---

## Gotchas
#### Is this code subject to any common programming gotchas?

Note:
Common gotchas: transposition errors, off by one errors, memory leaks, null dereferences
---
## Gotchas
#### Is spelling correct and consistent?

Note:
Amazon S3 Java library
---
## Language specific

#### Is the code idiomatic?
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

Note:
being able to teach is an opportunity. Don't waste it.
---

# Thank you

