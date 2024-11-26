const swaggerJsdoc = require('swagger-jsdoc');
const { SERVICE_PORT, SERVICE_NAME, SWAGGER_VERSION } = require('./constants');

const swaggerOptions = {
  swaggerDefinition: {
    swagger: '2.0',
    info: {
      title: `${SERVICE_NAME} API`,
      version: SWAGGER_VERSION,
      description: `Dynamic Swagger generation for ${SERVICE_NAME}`,
    },
    host: `localhost:${SERVICE_PORT}`,
    basePath: '/',
    schemes: ['http'],
  },
  apis: ['./routes/*.js'], // Path to API routes
};

const swaggerSpec = swaggerJsdoc(swaggerOptions);

module.exports = { swaggerSpec };
