# Temperaas

Template rendering as a service.

## What does it do

```
http://temperaas.alhur.es/http://somewhere.com/a-mustache-template.html?title1=something&title2=nothing
```

If `http://somewhere.com/a-mustache-template.html` holds a file like
```html
<template>
  <h1>{{ title1 }}</h1>
  <p>...</p>

  <h1>{{ title2 }}</h1>
  <p>...</p>
</template>
```

the output will be

```html
<template>
  <h1>something</h1>
  <p>...</p>

  <h1>nothing</h1>
  <p>...</p>
</template>
```
