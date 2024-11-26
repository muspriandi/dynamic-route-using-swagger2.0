const express = require('express');
const router = express.Router();

/**
 * @swagger
 * /member:
 *   post:
 *     summary: Create a new member
 *     parameters:
 *       - in: body
 *         name: body
 *         description: Member object to create
 *         required: true
 *         schema:
 *           type: object
 *           properties:
 *             name:
 *               type: string
 *               example: John Doe
 *             age:
 *               type: integer
 *               example: 30
 *     responses:
 *       201:
 *         description: Member created
 *         schema:
 *           type: object
 *           properties:
 *             message:
 *               type: string
 *             data:
 *               type: object
 */
router.post('/', (req, res) => {
  res.status(201).json({ message: 'Member created', data: req.body });
});

/**
 * @swagger
 * /member:
 *   get:
 *     summary: Get all members
 *     responses:
 *       200:
 *         description: A list of members
 *         schema:
 *           type: array
 *           items:
 *             type: object
 *             properties:
 *               id:
 *                 type: string
 *               name:
 *                 type: string
 *               age:
 *                 type: integer
 */
router.get('/', (req, res) => {
  res.status(200).json([
    { id: '1', name: 'John Doe', age: 30 },
    { id: '2', name: 'Jane Smith', age: 25 },
  ]);
});
/**
 * @swagger
 * /member/{id}:
 *   put:
 *     summary: Update a member by ID
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         type: string
 *         description: The ID of the member to update
 *       - in: body
 *         name: body
 *         description: Member object with updated details
 *         required: true
 *         schema:
 *           type: object
 *           properties:
 *             name:
 *               type: string
 *             age:
 *               type: integer
 *     responses:
 *       200:
 *         description: Member updated
 *         schema:
 *           type: object
 *           properties:
 *             message:
 *               type: string
 *             data:
 *               type: object
 */
router.put('/:id', (req, res) => {
  const { id } = req.params;
  res.status(200).json({ message: `Member ${id} updated`, data: req.body });
});

/**
 * @swagger
 * /member/{id}:
 *   delete:
 *     summary: Delete a member by ID
 *     parameters:
 *       - in: path
 *         name: id
 *         required: true
 *         type: string
 *         description: The ID of the member to delete
 *     responses:
 *       200:
 *         description: Member deleted
 *         schema:
 *           type: object
 *           properties:
 *             message:
 *               type: string
 */
router.delete('/:id', (req, res) => {
  const { id } = req.params;
  res.status(200).json({ message: `Member ${id} deleted` });
});

module.exports = router;