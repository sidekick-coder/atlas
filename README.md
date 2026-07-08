## todo
- [x] tui: query filter in the list screen
- [x] tui: entry single page defulat page for entry with a list of metas
- [x] tui: sync page
- [x] sync all multithreaded
- [x] bug in single page not able to edit metas
- [ ] load metadata handlers from config
- [ ] custom metadata handler (probably comunicate via stdin/stdout)
- [ ] tui: move queries data to components instead of the screen
- [ ] tui: columns in the list screen

## ideas 

- tui: keymaps per filetype using groups 
- - ext:md  for entry markdown files
- - hasmeta:frontmatter.status if entry has a meta key called frontmatter.status
- - meta:frontmatter.status when meta is available, eg editin a meta
- tui: customize what screen to use for each entry type, for example a markdown file can be opened in a single markdown screen or in a single text screen.
- tui: manualy open a screen and pass its params
- entry single text: screen to view text files with a metas sidebar and a body
- entry single markdown: screen to view markdown files with a frontmatter sidebar and a body
- entry kanban: view entries in a kanban board, with columns based on a meta key
- user predefined screens: user can define screens with a set of filters and a layout, for example a kanban board with columns based on a meta key, or a list of entries with a sidebar showing the metas of the selected entry.
- components: elements that the user can use to compose a screen, entry list, kanba, chart, etc.
- - custom componets are sheel scripts than recieve width and height and print the text to show in the screen
- - buint-int components can be used too like, meta field input
- - the components will have its own keymaps that are active when the component is focused
- screens: screens are a list of components 
- - define x,y,width,height and render in the screen area
- - scroll support for each component

- action handlers: type of actions to execute actions in the app 
- - toast: show a toast message 
- - open-screen: open a screen with a set of options
- - action-group: execute a set of actions in order


## Open 

You can defined witch command use to open an entry, for example you can use `code` to open a file in VSCode or `nvim` to open it in Neovim.

```toml 
open.handlers.markdown.pattern=*.md
open.handlers.markdown.command="code {{entry.path}}"
```

Or integrate with tmux for example:

```toml
open.handlers.markdown.pattern=*.md
open.handlers.markdown.command="tmux new-window -n {{entry.name}} 'nvim {{entry.path}}'"
```

## Sync 

To query entries correctly, the tool needs to have a local database that is synced with the workspace.

The first time you use the toll you will be required to sync all entries available.

```sh
atlas sync-all
```

This need to happen only once, and then every time you modify something via the tool, the database will be updated automatically.

If you modify something outside the tool, you don't need to sync all again you can make a focus sync with run `atlas sync`.

```sh 
atlas sync tasks/001.md
```

## Search entries 

You can query entries using a special syntax, for example:

```sh
atlas list -q "type=file ext=md frontmatter.id=001"
```

This returns all entries that are files, have the extension `.md`, and have a frontmatter tag of `project`.

## Filters 

### type 

Filter by entry type, for example `type=file` or `type=directory`.

### ext 

Filter by file extension, for example `ext=md` or `ext=txt`.

### frontmatter 

This is a filter that exists for markdown files that have frontmatter. You can filter by frontmatter tags, for example `frontmatter.tag=project` or `frontmatter.status=active`.

### Others 

Basic you can query any metadata that is stored in the entry_metas table, for example `myspecialmeta=somevalue`.

## JSON output 

You can output the results in JSON format using the `-j` flag:

```sh
atlas list -q "type=file ext=md frontmatter.tag=project" -j
```

## entities 

Define some entities that live in the workspace, normally it is a list of files in folder or something like that 

And entity have a glob pattern that matches the entries in the workspace.

Also the entities have a set of fields that map to entry metas

## Metadata handler 

This are scripts to handle extraction and updates of metadata from entries. 

For example the markdown handler is responsible for extracting frontmatter from markdown files and saving it in the entry_metas table.

The metas are store this way: 

```tom
frontmatter.id=00 
frontmatter.status=active 
frontmatter.tags[0]=project
frontmatter.tags[1]=important
```

The metas follow a dot notation, and arrays are represented with square brackets.

This is to help with the search queries, for example you can query `frontmatter.tags[0]=project` or `frontmatter.tags[1]=important`.
