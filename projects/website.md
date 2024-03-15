---
id: 1708763107-HFJZ
aliases:
  - personal-website
tags:
  - go
  - htmx
  - obsidian
area: ""
project: ""
---

# personal-website

### source code: [click](https://github.com/nynniaw12/go-htmx-aybars)

This website uses the very new tech stack of Go and HTMX. It sports complete
server-side rendering and lives on a Docker container. All of the styling is done
using TailwindCSS. It is still under development with a lot of rough edges.

It will sync with git to fetch obsidian markdown notes on my projects, almost like
development logs. It currently uses goldmark written in pure Go as its markdown 
renderer, however, there is integration problems with TailwindCSS as can be observed 
from the currently poor readability.

## TODO:

- [x] Dockerize website
- [ ] Host on own server
- [ ] Sync with git on the server 
    - [ ] Add a bind mount to projects folder
- [ ] Mailing for contact
- [ ] Complete UI
    - [x] Go md renderer
        - [x] Serving files and preprocessing md files
        - [ ] Fix spacing and improve readability
    - [ ] Carousel with HTMX
    - [x] Fix tabbing with HTMX code snippet

