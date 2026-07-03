- [ ] search entries
- [ ] global db connection and config
- [ ] plan how extract metadata from entries (markdown frontmatter)
- [ ] plan how to handle custom actions (create-task, create-project, etc.)

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
