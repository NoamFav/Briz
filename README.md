# 🌬 Briz

<div align="center">

<img src="https://img.shields.io/badge/go-1.25+-00ADD8.svg?style=for-the-badge&logo=go" alt="Go">
<img src="https://img.shields.io/badge/license-MIT-green.svg?style=for-the-badge" alt="License">

**AI-powered file and directory organizer**

[Installation](#installation) · [Quick Start](#quick-start) · [Usage](#usage)

</div>

---

Briz uses an LLM to analyze your directory tree and automatically sort files into a logical structure — no manual rules needed.

```
briz <path>    → analyze and sort the directory at <path>
```

---

## Installation

```bash
git clone https://github.com/NoamFav/Briz
cd Briz && make install
```

---

## Usage

```bash
# Sort a directory using AI suggestions
briz ~/Downloads

# Preview changes without applying
briz ~/Downloads --dry-run
```

---

## License

MIT — see [LICENSE](LICENSE).

---

<div align="center">
Made with ❤️ by <a href="https://github.com/NoamFav">NoamFav</a>
</div>
