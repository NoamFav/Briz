<div align="center">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/brand/banner-night.svg">
  <img alt="Briz: AI-powered CLI that reads your directory tree and reorganizes it into a logical structure. Dry-run first, no manual rules." src=".github/brand/banner-paper.svg" width="100%">
</picture>
<br><br>
<a href="#about"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-about-night.svg"><img alt="about" src=".github/brand/tab-about-paper.svg"></picture></a>
<a href="#how-it-works"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-how-it-works-night.svg"><img alt="how it works" src=".github/brand/tab-how-it-works-paper.svg"></picture></a>
<a href="#quickstart"><picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/tab-quickstart-night.svg"><img alt="quickstart" src=".github/brand/tab-quickstart-paper.svg"></picture></a>
</div>

<p>
<a name="about"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-about-night.svg"><img alt="about" src=".github/brand/section-about-paper.svg" width="100%"></picture>
</p>

Briz points an LLM at a messy directory, has it read the file tree, and sorts everything into a logical structure — no hand-written rules, no regex matching on filenames.

```
briz ~/Downloads              → analyze and sort
briz ~/Downloads --dry-run    → preview the plan without touching anything
```

<p>
<a name="how-it-works"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-how-it-works-night.svg"><img alt="how it works" src=".github/brand/section-how-it-works-paper.svg" width="100%"></picture>
</p>

| Package | Role |
|---------|------|
| `internal/fileops` | Filesystem walking, moves, safety checks |
| `internal/llm` | Prompts an LLM with the directory tree, parses the sort plan |
| `internal/rules` | Loads `config/default_rules.yaml` as guardrails for the LLM |
| `internal/sorter` | Applies the plan — or just prints it under `--dry-run` |

<p>
<a name="quickstart"></a>
<picture><source media="(prefers-color-scheme: dark)" srcset=".github/brand/section-quickstart-night.svg"><img alt="quickstart" src=".github/brand/section-quickstart-paper.svg" width="100%"></picture>
</p>

```sh
git clone https://github.com/NoamFav/Briz && cd Briz
make install

briz ~/Downloads --dry-run   # see the plan
briz ~/Downloads             # apply it
```

<div align="center">

Made with ♥ by [NoamFav](https://github.com/NoamFav) · Apache 2.0

</div>

<br>

<a href="https://nf-software.com">
<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/brand/footer-night.svg">
  <img alt="NF Software" src=".github/brand/footer-paper.svg" width="100%">
</picture>
</a>
