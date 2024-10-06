# Raindrop Images Downloader

> A CLI command to download images from [Raindrop](https://raindrop.io) Collections.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/brpaz/raindrop-images-dl?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/brpaz/raindrop-images-dl?style=for-the-badge)](https://goreportcard.com/report/github.com/brpaz/raindrop-images-dl)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/brpaz/raindrop-images-dl/ci.yml?style=for-the-badge)](https://github.com/brpaz/raindrop-images-dl/actions)
[![Codecov](https://img.shields.io/codecov/c/github/brpaz/raindrop-images-dl?style=for-the-badge)](https://app.codecov.io/gh/brpaz/raindrop-images-dl)
[![GitHub License](https://img.shields.io/github/license/brpaz/raindrop-images-dl?style=for-the-badge)](LICENSE)

## 🏹 Motivation

I am using Raindrop to archive Internet images like Gifs and Memes.

"Permanent Copy" helps to retain images that might be deleted, but there is no simple way to have a structure backup from this data in Raindrop.

Unfortunately the export feature of Raindrop is limited regarding images and uploaded files, so I built this tool to help with the backup process.


## 🚀 Getting started

### Pre-requisites

- A [Raindrop](https://raindrop.io/) account.
- A Raindrop API **Test token** - Check [here](https://developer.raindrop.io/v1/authentication/token)

### Install

This application is built as a single binary so the installation is very easy. Just heads up to the [Releases pages](https://github.com/brpaz/raindrop-images-dl/releases) and download the latest version for your Operating system.

You can also use Docker (and NixOS support is planned).

### Usage

Get the `collection id` of the Raindrop collection you want to download from. You can find it in the url, when accessing Raindrop Web App. Ex: `https://app.raindrop.io/my/<collection_id>`.

Then run the following command:

```shell
raindrop-images-dl download \
    -c=<my_collection_id>
    -k <raindrop_api_key>
    -o <path/to/images/dir>
```

Alternatively, envrionment variables can also be used instead of flags:

```shell
RAINDROP_API_KEY=<api_key> RAINDROP_COLLECTION=<collection_id> OUTPUT_DIR=<output> raindrop-images-dl download
```

The command will download the images found in your collection, to the output directory defined with `-o` option.

A subfolder with the collection name, will be created.

A `.info.json` file will be placed together with the image file. This file will save some Raindrop metadata like tags.

You can use this for doing some automations.

## 🤝 Contributing

All contributions are welcome. Please see [CONTRIBUTING.md](CONTRIBUTING.md) file for details.

## 🫶 Support

If you find this project helpful and would like to support its development, there are a few ways you can contribute:

[![Sponsor me on GitHub](https://img.shields.io/badge/Sponsor-%E2%9D%A4-%23db61a2.svg?&logo=github&logoColor=red&&style=for-the-badge&labelColor=white)](https://github.com/sponsors/brpaz)

<a href="https://www.buymeacoffee.com/Z1Bu6asGV" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## 📩 Contact

✉️ **Email** - [oss@brunopaz.dev](oss@brunopaz.dev)

🖇️ **Source code**: [https://github.com/brpaz/raindrop-images-dl](https://github.com/brpaz/raindrop-images-dl)

## 📃 License

Distributed under the MIT License.
See [LICENSE](LICENSE) file for details.
