<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
<!--   <a href="https://github.com/tarumes/gotypo">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">gotypo</h3>

  <p align="center">
    a levenstein based word replacer with proposals
    <br />
    <a href="https://github.com/tarumes/gotypo"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!-- <a href="https://github.com/tarumes/gotypo">View Demo</a> -->
    ·
    <a href="https://github.com/tarumes/gotypo/issues">Report Bug</a>
    ·
    <a href="https://github.com/tarumes/gotypo/issues">Request Feature</a>
    ·
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

a levenstein based word replacer with proposals

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

[![Golang][golang]][golang-url]


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* go get 
  ```sh
  go get -u https://github.com/tarumes/gotypo@master
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

```Go
package main

import (
	"log"

	"github.com/tarumes/gotypo"
)

func main() {

	corrector := gotypo.New()
	if err := corrector.LearnWord([]string{"Hello", "World"}); err != nil {
		log.Fatal(err)
	}
	sentence := corrector.TypoCorrections("Hewo Wold")
	log.Println(sentence)

}
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] single word correction
- [x] sentence correction
- [ ] syntax correction

See the [open issues](https://github.com/tarumes/gotypo/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>







<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/tarumes/gotypo.svg?style=for-the-badge
[contributors-url]: https://github.com/tarumes/gotypo/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/tarumes/gotypo.svg?style=for-the-badge
[forks-url]: https://github.com/tarumes/gotypo/network/members
[stars-shield]: https://img.shields.io/github/stars/tarumes/gotypo.svg?style=for-the-badge
[stars-url]: https://github.com/tarumes/gotypo/stargazers
[issues-shield]: https://img.shields.io/github/issues/tarumes/gotypo.svg?style=for-the-badge
[issues-url]: https://github.com/tarumes/gotypo/issues
[product-screenshot]: images/screenshot.png
[golang]: https://img.shields.io/github/go-mod/go-version/tarumes/gotypo
[golang-url]: https://go.dev/
