# Raindrop Alfred Workflow

[![Actions Status](https://github.com/ytakahashi/raindrop-alfred-workflow/workflows/Go%20CI/badge.svg)](https://github.com/ytakahashi/raindrop-alfred-workflow/actions)
[![GitHub release](https://img.shields.io/github/release/ytakahashi/raindrop-alfred-workflow.svg)](https://github.com/ytakahashi/raindrop-alfred-workflow/releases/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Alfred workflow to see bookmarks stored on [Raindrop.io](https://raindrop.io/).

This workflow utilizes [Raindrop.io API](https://developer.raindrop.io).

## How to use

Download a latest package from [release page](https://github.com/ytakahashi/raindrop-alfred-workflow/releases) and install.

When installing this workflow, variable named "accessToken" should be configured.

You can obtain your access token to follow [this document](https://developer.raindrop.io/v1/authentication/token). I recommend to use test token because this workflow does not access any data except your account.

Activate Alfred and `⇧⌘R` triggers this workflow.  

![workflow image](./image/workflow_image_1.png)

`Raindrops` calls [Get raindrops API](https://developer.raindrop.io/v1/raindrops/multiple#get-raindrops).

`Collections` calls [Get root collections API](https://developer.raindrop.io/v1/collections/methods#get-root-collections).
