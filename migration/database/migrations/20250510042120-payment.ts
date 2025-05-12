import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('payment')
    if (!tablExist) {
      await queryInterface.createTable(
        'payment',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          bank_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'bank', key: 'id' } },
          user_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'users', key: 'id' } },
          request_id: { type: DataTypes.UUID, unique: true, allowNull: false },
          amount: { type: DataTypes.BIGINT },
          status: { type: DataTypes.STRING(25), allowNull: false },
          sender: { type: DataTypes.STRING(200) },
          account_number: { type: DataTypes.INTEGER({ unsigned: true }) },
          verified_at: { type: DataTypes.DATE },
          expired_at: { type: DataTypes.DATE },
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
        queryInterface.addIndex('payment', ['id', 'bank_id', 'user_id', 'request_id'], { name: 'payment_id_bank_id_user_id_request_id_deleted_at_idx', unique: true, logging: true }),
        queryInterface.addIndex('payment', ['status', 'verified_at', 'deleted_at'], { name: 'payment_status_verified_at_deleted_at_idx', unique: true, logging: true }),
      ])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('payment')
    if (tableExist) {
      return queryInterface.dropTable('payment', { logging: true })
    }
  },
}
