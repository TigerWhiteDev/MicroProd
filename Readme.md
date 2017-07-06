# Rinyx Framwork

## List of functionality

```
-  all configs files on YAML
-  frontend and backend separated
-  possible to use anyware databases  ( module core ORM)
-  possible to send the services  on different server
-  possible to dev service with different language 
-  possible to realise a frontend on mobile application , desktop software 
-  possible to deploy automatiqueli on docker and docker-swarm
-  gestion of dependancy for service
-  gestion of unitarie test 

```
## Specification

### Command
```
rinyx install ( install all depency of framwork and services)
rinyx generate application -language lang -backend | -frontend
rinyx generate module application -backend | -frontend
rinyx build (build all service)
rinyx start ( run mode dev)
rinyx stop ( stop mode dev)
rinyx deploy -targz|-docker ( generate  dockerfile and docker-compose.yml or tar.gz )


```

### call librairy

```
rinyx.model.NameEntity
rinyx.settings.NameSettingsfile
rinyx.storage.NameStorage

```