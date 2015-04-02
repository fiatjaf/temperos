# Temperaas

Template rendering as a service.

## What does it do

If `http://somewhere.com/a-mustache-template.html` holds a file like
```html
<template>
  <h1>{{ title1 }}</h1>
  <p>...</p>

  <h1>{{ title2 }}</h1>
  <p>...</p>
</template>
```

the output of

```
http://temperaas.alhur.es/http://somewhere.com/a-mustache-template.html
  ?title1=something&title2=nothing
```

will be

```html
<template>
  <h1>something</h1>
  <p>...</p>

  <h1>nothing</h1>
  <p>...</p>
</template>
```

## Examples

* Get [this random gist URL](https://gist.githubusercontent.com/nmerouze/258889/raw/e19df1014ce205fbc82dabe7edd6cdb2b2b7c71b/index.html), add some parameters and it renders into [this nice question](http://temperaas.alhur.es/https://cdn.rawgit.com/nmerouze/258889/raw/e19df1014ce205fbc82dabe7edd6cdb2b2b7c71b/index.html?title=Who&calc=that%20much?).
* Or perhaps there's something more useful than this: you can get a [Google Analytics empty snippet](https://gist.githubusercontent.com/fiatjaf/24aee0052afc73035ee6/raw/e4060e9348079792a098d42fa8ad8b3c2bf2aee5/add-google-analytics.js) and render it into your [custom snippet with the correct code](http://temperaas.alhur.es/https://cdn.rawgit.com/fiatjaf/24aee0052afc73035ee6/raw/e4060e9348079792a098d42fa8ad8b3c2bf2aee5/add-google-analytics.js?code=YOUR_GOOGLE_ANALYTICS_TRACKING_CODE).
* Can you think of more use cases? Please add them.
