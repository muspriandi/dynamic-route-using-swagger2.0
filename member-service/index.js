const express = require('express');
const path = require('path');
const fs = require('fs');
const swaggerUi = require('swagger-ui-express');
const { swaggerSpec } = require('./lib/swagger');
const { SERVICE_PORT, SERVICE_CODE, SERVICE_NAME } = require('./lib/constants');

// Initialize the app
const app = express();
app.use(express.json());

// Import routes
const pingRoutes = require('./routes/ping');
const memberRoutes = require('./routes/member');

// Use routes
app.use('/ping', pingRoutes);
app.use('/member', memberRoutes);

// Serve Swagger documentation
app.use('/swagger', swaggerUi.serve, swaggerUi.setup(swaggerSpec));

// Export Swagger JSON
const swaggerFilePath = path.join(__dirname, `swagger/${SERVICE_CODE}-swagger-spec.json`);
fs.writeFileSync(swaggerFilePath, JSON.stringify(swaggerSpec, null, 2));
console.log(`Swagger spec written to ${swaggerFilePath}`);

// Start the server
app.listen(SERVICE_PORT, () => {
  console.log(`${SERVICE_NAME} server running on http://localhost:${SERVICE_PORT}`);
});