import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('order')
    if (!tablExist) {
      await queryInterface.createTable(
        'order',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          payment_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'payment', key: 'id' } },
          user_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'users', key: 'id' } },
          courier_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'courier', key: 'id' } },
          invoice_number: { type: DataTypes.STRING(50), unique: true, allowNull: false },
          status: { type: DataTypes.STRING(25), allowNull: false, defaultValue: 'pending' },
          paid: { type: DataTypes.BOOLEAN, allowNull: false, defaultValue: false },
          notes: { type: DataTypes.TEXT },
          origin_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
          total_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
          discount_amount: { type: DataTypes.INTEGER({ unsigned: true }), defaultValue: 0 },
          created_at: { type: DataTypes.DATE, allowNull: false, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
          created_by: { type: DataTypes.STRING(200) },
          updated_at: { type: DataTypes.DATE },
          updated_by: { type: DataTypes.STRING(200) },
          deleted_at: { type: DataTypes.DATE },
          deleted_by: { type: DataTypes.STRING(200) },
        },
        {
          logging: true,
        },
      )

      return Promise.all([
        queryInterface.addIndex('order', ['id', 'payment_id', 'user_id', 'deleted_at'], { name: 'order_id_payment_id_user_id_deleted_at_idx', unique: true, logging: true }),
        queryInterface.addIndex('order', ['invoice_number', 'status', 'paid'], { name: 'order_invoice_number_status_paid_idx', unique: true, logging: true }),
      ])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('order')
    if (tableExist) {
      return queryInterface.dropTable('order', { logging: true })
    }
  },
}
