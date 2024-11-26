const express = require('express');
const router = express.Router();


/**
 * @swagger
 * /ping:
 *   get:
 *     summary: Responds with pong
 *     responses:
 *       200:
 *         description: A successful response
 *         schema:
 *           type: object
 *           properties:
 *             message:
 *               type: string
 */
router.get('/', (req, res) => {
  res.status(200).json({ message: 'pong order' });
});

module.exports = router;