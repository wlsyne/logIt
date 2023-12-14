# The user stories for the Changelog writing tool

## User Story 1: Users can configure the changelog writing tool with some parameters
- title: string, the user can configure the changelog title with the parameter `title`
- gitBaseUrl: string,the user can configure the git base url with the parameter `gitBaseUrl`,then the tool can generate the git commit url with adding the commit id to the git base url 
- chatIds: string[], the user can configure the chat ids with the parameter `chatIds`
- botWebhook: string, the user can configure the bot webhook with the parameter `botWebhook`

## User Story 2: Users can write different types of changelog 
- When user type `logIt write` or `logIt w`, the tool will enter the write mode
- Then the tool will ask the user to select the type of the changelog
  - Feat: A new feature
  - Docs: Documentation only changes
  - Fix: A bug fix
  - Style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
  - SpeedUp: Changes that improve performance
  - Config: Changes to config files
  - Test: Adding missing tests or correcting existing tests
  - BreakChange: A breaking change
  - Finish: Finish writing the changelog
  - Cancel: Cancel writing the changelog
  - Finish: Finish writing the changelog
- Then the tool will ask the user to select the commit id from the git log
- If the user select `Finish` or `Cancel`, the tool will exit the write mode
  - If the user select `Cancel`, the tool will not write the written changelog to the changelog file and exit the write mode
  - If the user select `Finish`, the tool will write the written changelog to the changelog file  - If the user select `Finish`, the tool will write the written changelog to the changelog file
- When the user select the Finish option
  - title: the tool will pick the title from config file ,if the user has not configured the title, the tool will use the default title `name + version` which comes from `package.json`
  - userName: the tool will pick the userName from git
- Then the tool will write the completed changelog to the changelog file

### changelog example
```markdown
# @a/b 1.2.2
- ✨ Feat: log间距改为两行  [#8ab0e0e](https://a/b/c/commit/8ab0e0e)
> Published by <@synwu>
```

## User Story 3: Users can publish the recent changelog to the WeCom group
- When user type `logIt publish` or `logIt p`, the tool will enter the publishing mode
- The tool will pick the most recent changelog from the changelog file to publish
- Then the tool will pick the chatIds and botWebhook from config file, if the user has not configured the chatIds and botWebhook, the program will remind the user to configure the chatIds and botWebhook and exit the publishing mode
- Then the tool will publish the changelog to the WeCom group through the bot SDK of WeCom