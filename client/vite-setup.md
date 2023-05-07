# Steps to Setup Vite

1. `npm create vite@latest`
2. cd <project_name_in_step_1>
3. `npm install`
4. `npm run dev`
5. Install tailwind css
   - `npm install -D tailwindcss postcss autoprefixer`
   - `npm tailwindcss init -p`
6. Add to index.css
   ```css
   @tailwind base;
   @tailwind components;
   @tailwind utilities;
   ```
7. Change port in Vite
   - `vite.config.ts`
   ```js
   export default defineConfig({
     plugins: [react()],
     server: {
       port: 3000,
     },
   });
   ```
