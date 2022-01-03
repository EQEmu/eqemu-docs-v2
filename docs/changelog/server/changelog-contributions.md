!!! info
    
    This page outlines notes regarding maintaining the changelog

### Guidelines

Your changelog should describe the end user functionality and why it is important. Similar thing goes for submitting pull requests.

While listing PR titles in the changelog is not preferred; sometimes our community gets so busy that it's all that we have time to put up in the changelog.

If you can, try to put more detailed notes in the server changelog. The purpose of the changelog is to inform users of the EverQuest Emulator Server software of what is new, what has changed, what's better, what's fixed and that is not always surfaced with obscure PR messages.

### Git Log Message Generate

!!! Akkadius

    This is what I typically do when I'm doing bulk changelogs. If someone else ever wants to do it.

If you do choose to generate the changelog using at least the PR merges use the following command in a bash shell (git bash on windows) or otherwise. Adjust your days back for how far back you want to go

```
git log --no-merges --first-parent --since=40.days --pretty=format:'**%an**%Creset%C(yellow)%d%Creset %s' --abbrev-commit | sort
```

Take that output and bring it into a Github new issue under EQEmu/Server; we're doing this simply to generate the pull request links

![image](https://user-images.githubusercontent.com/3319450/147895791-c4d1f241-8581-4884-8591-e65dedbfd922.png)

Once you have the content inserted, go to the preview pane, open up inspect element in the browser and we're going to copy the HTML. 

![image](https://user-images.githubusercontent.com/3319450/147895889-6e67fd9b-bd96-4212-8f88-6b02b155c9b6.png)

From the HTML we need convert back to markdown. Use something like [https://www.browserling.com/tools/html-to-markdown](https://www.browserling.com/tools/html-to-markdown)

There you go, paste the markdown into a header section containing the date and the entries below.

![image](https://user-images.githubusercontent.com/3319450/147895976-9a4baae8-6211-49a2-a0ea-a481e0c689bd.png)
