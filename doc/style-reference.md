# Style reference

~~~html
<style>
.container {
   background: green;
   display: flex;
   flex-wrap: wrap;
   gap: 10px;
   justify-content: center;
   margin: 10px auto;
   max-width: 700px; /* fix */
}
.item {
   background: red;
   height: 100px;
   width: 200px;
}
</style>
<div class="container">
   <div class="item"></div>
   <div class="item"></div>
</div>
<div class="container">
   <div class="item"></div>
   <div class="item"></div>
   <div class="item"></div>
   <div class="item"></div>
</div>
~~~

~~~html
<style>
.container {
   background: green;
   display: grid;
   gap: 10px;
   grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
   justify-items: center;
   margin: 10px auto;
   max-width: 800px; /* fix */
}
.item {
   background: red;
   height: 100px;
   width: 100px;
}
</style>
<div class="container">
   <div class="item"></div>
   <div class="item"></div>
</div>
<div class="container">
   <div class="item"></div>
   <div class="item"></div>
   <div class="item"></div>
   <div class="item"></div>
</div>
~~~

**Color**:

- <https://color.adobe.com>
- <https://colorhunt.co>
- <https://colorkoala.netlify.app>
- <https://tachyons.io>
- <https://tailwindcss.com/docs/customizing-colors>

**Data types**:

Avoid `ex` as it produces inconsistent results on mobile browsers. Use `em`
instead.
