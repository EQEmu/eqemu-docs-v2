Looking to add to the documentation? Not sure how? You've found the right place

In this page we'll go over how to add new pages, formatting, tips and tricks, guidelines etc.

## Markdown

First of all, if you're not familiar with markdown - every single page in the docs are written in [markdown](https://www.markdownguide.org/getting-started/). Check out the markdown getting started page and get familiar with it.

## MkDocs Material

While there is standard markdown syntax, we have a variety of **markdown extensions** that we use from our documentation platform [MkDocs Material](https://squidfunk.github.io/mkdocs-material/reference/abbreviations/) of which you can reference for different ways to format different elements

## Making Small Page Edits

You will see on most pages in the upper right corner of the page there is an **Edit This Page** which will bring you straight to Github

When you click edit, it will bring you to Github

![image](https://user-images.githubusercontent.com/3319450/143397077-3ece1596-3dad-49af-8c62-51a3b3315f4f.png)

You can click the **Preview** tab to help you see what your markdown looks like

![image](https://user-images.githubusercontent.com/3319450/143397182-b1b54f18-d6b7-45ce-b45b-110cd8309c2f.png)

Notice `!!! info` is not formatted in the Github preview, that is because it is custom MkDocs markdown formatting. Using this method works for light editing but if you plan on writing documentation regularly you should install MkDocs

You can submit a pull request or a commit to this repository and changes when approved and merged will automatically be published to this site.

## Install MkDocs Material

Clone our documentation repository 

```text
git clone https://github.com/EQEmu/eqemu-docs-v2.git
cd eqemu-docs-v2
```

[Install mkdocs](https://squidfunk.github.io/mkdocs-material/getting-started/)

Requires having Python installed

```text
pip install mkdocs-material
```

## Run the Documentation Server

Once it is successfully installed you can run the development docs server

```text
mkdocs serve --dirtyreload

INFO     -  Documentation built in 15.81 seconds
INFO     -  [01:26:45] Serving on http://127.0.0.1:8000/
INFO     -  [01:26:46] Browser connected: http://127.0.0.1:8000/contributing/how-to-doc/
INFO     -  [01:27:13] Detected file changes
```

We use the `--dirtyreload` for speed reasons and it's very important. Every time we make a change to a simple page we don't want to rebuild the whole entire thing. We have well over 600 pages of documentation and this would take too long.

The only time you need to re-run this command is when you make navigation changes to `mkdocs.yml`. Small changes to markdown files don't require re-running `mkdocs serve`

![image](https://user-images.githubusercontent.com/3319450/143397894-e593f869-de9b-4ebc-a705-b1d305b16a5c.png)

## Editing Pages

Now that you have the MkDocs Material Documentation Server running, it's a good idea to use some sort of text editor locally to edit the files. 

You can use a simple text editor to edit markdown files. If you want a visual aid for what your markdown looks like you can install a markdown extension in your editor of choice like Visual Studio Code, Sublime, Jetbrains IDE's etc.

![image](https://code.visualstudio.com/assets/docs/languages/Markdown/preview-scroll-sync.gif)


## Uploading Images

Very often you will want to use images to illustrate points in your docs. How do you upload an image? Where do you put it? Well, there are a few options.

You can upload images via commit git but its not recommended to be storing images at scale in the repository. You also want to make sure that if you are using an image that it is uploaded to a reliable image host so we don't have dead image links in the future.

The most reliable mechanism for uploading images is to use a hack that I use which is to use the Github issues page to copy and paste images into the text area and it will return a markdown code for you to use in your page.

![image](https://i.stack.imgur.com/RKu7x.gif)

## Adding a New Page to the Navigation Tree

The documentation navigation tree configuration is stored in `mkdocs.yml` at the top of the documentation repository in the `nav` section.

[https://github.com/EQEmu/eqemu-docs-v2/blob/main/mkdocs.yml](https://github.com/EQEmu/eqemu-docs-v2/blob/main/mkdocs.yml)

Do not nest pages more than 1-2 levels deep. We should not have to gopher dig for information and should be easily accessible within a simple 1-2 categorization method.

For example if your new page has something to do with inventory logic, then it should be under a top level category called **Inventory** not something archaic like **Developer Notes** -> **Developer Section** -> **Inventory Notes** -> **Inventory Logic**

