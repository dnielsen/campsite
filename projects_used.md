Frontend
- NodeJS
- React
- Material UI
- Next.js
- JSON Server (db.json) https://github.com/typicode/json-server/
- ESLint

Backend
- Postgres
- 
How the DB.JSON version works:
Browser requests localhost:3000 which is served by NextJS
- which routes the request to index.tsx 
- index.tsx needs some data so NextJS makes 3 requests to the rest api provided by json server:
 localhost:4444/events/id -> reads data in db.json
 localhost:4444/sessions -> reads data in db.json
 localhost:4444/speakers -> reads data in db.json
Nextjs combines the results of the API calls with the React component ../components/event/EventItem.tsx 
Outputs the results to the web browser. Each page has it's own URL, and helps improve SEO.

How the Monolith version works:
Browser requests localhost:3000 which is served by NextJS
- which routes the request to index.tsx 
- index.tsx needs some data so NextJS makes 1 requests to the rest api provided by json server:
 localhost:4444/events/id -> reads Event, Sessions & Speaker data from Postgres
Nextjs combines the results of the API call with the React component ../components/event/EventItem.tsx
Outputs the results to the web browser. Each page has it's own URL, and helps improve SEO.
