<div align="center">

<img src="https://capsule-render.vercel.app/api?type=venom&height=220&color=gradient&customColorList=12&text=BRIZ&fontSize=100&fontColor=fff&animation=twinkling&desc=An%20LLM%20Cleans%20Up%20Your%20Downloads%20Folder&descSize=18&descAlignY=65&stroke=FFFFFF&strokeWidth=1" alt="Briz Banner" />

<img src="https://readme-typing-svg.herokuapp.com?font=Fira+Code&size=20&pause=1000&color=00D9FF&center=true&vCenter=true&multiline=true&repeat=true&width=900&height=60&lines=briz+~%2FDownloads+%E2%86%92+an+LLM+reads+the+tree%2C+sorts+it;No+manual+rules+%C2%B7+dry-run+first+%C2%B7+Go" alt="Typing SVG" />

<br>

[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go&logoColor=white&labelColor=0D1117)](https://go.dev)
[![License](https://img.shields.io/badge/MIT-00D9FF?style=for-the-badge&labelColor=0D1117)](./LICENSE)

</div>

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=00D9FF&center=true&width=800&lines=%F0%9F%A4%96+WHAT+IS+BRIZ+%3F" alt="What is Briz" />
</div>
<br>

Briz points an LLM at a messy directory, has it read the file tree, and sorts everything into a logical structure — no hand-written rules, no regex matching on filenames.

```
briz ~/Downloads              → analyze and sort
briz ~/Downloads --dry-run    → preview the plan without touching anything
```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=FF69B4&center=true&width=800&lines=%F0%9F%A7%A9+HOW+IT+WORKS+%F0%9F%A7%A9" alt="How it works" />
</div>
<br>

| Package | Role |
|---------|------|
| `internal/fileops` | Filesystem walking, moves, safety checks |
| `internal/llm` | Prompts an LLM with the directory tree, parses the sort plan |
| `internal/rules` | Loads `config/default_rules.yaml` as guardrails for the LLM |
| `internal/sorter` | Applies the plan — or just prints it under `--dry-run` |

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">
  <img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=26&pause=1000&color=6A5ACD&center=true&width=800&lines=%E2%9C%A8+QUICKSTART+%E2%9C%A8" alt="Quickstart" />
</div>
<br>

```sh
git clone https://github.com/NoamFav/Briz && cd Briz
make install

briz ~/Downloads --dry-run   # see the plan
briz ~/Downloads             # apply it
```

<img src="https://user-images.githubusercontent.com/73097560/115834477-dbab4500-a447-11eb-908a-139a6edaec5c.gif" width="100%">

<div align="center">

<img src="https://readme-typing-svg.herokuapp.com?font=Orbitron&size=20&pause=1000&color=6A5ACD&center=true&width=800&lines=Thanks+for+stopping+by!" alt="Footer typing" />

<br>

Made with ♥ by [NoamFav](https://github.com/NoamFav) · MIT License

<img src="https://capsule-render.vercel.app/api?type=waving&height=100&color=gradient&customColorList=12&section=footer" />

</div>
