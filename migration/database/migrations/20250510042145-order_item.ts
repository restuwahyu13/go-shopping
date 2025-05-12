import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
	up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
		const tablExist: boolean = await queryInterface.tableExists('order_item')
		if (!tablExist) {
			await queryInterface.createTable(
				'order_item',
				{
					id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
					order_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'order', key: 'id' } },
					product_item_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'product_item', key: 'id' } },
					qty: { type: DataTypes.INTEGER, allowNull: false },
					origin_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
					total_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
					discount_amount: { type: DataTypes.INTEGER, defaultValue: 0 },
					promotion_rules: { type: DataTypes.JSONB },
					free_product: { type: DataTypes.ARRAY(DataTypes.STRING(50)) },
					notes: { type: DataTypes.TEXT },
					created_at: { type: DataTypes.DATE, allowNull: false, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
					created_by: { type: DataTypes.STRING(200) },
					updated_at: { type: DataTypes.DATE },
					updated_by: { type: DataTypes.STRING(200) },
					deleted_at: { type: DataTypes.DATE },
					deleted_by: { type: DataTypes.STRING(200) }
				},
				{
					logging: true
				}
			)

			return Promise.all([
				queryInterface.addIndex('order_item', ['id', 'order_id', 'product_item_id', 'deleted_at'], { name: 'order_id_order_id_product_item_id_deleted_at_idx', unique: true, logging: true })
			])
		}
	},
	down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
		const tableExist: boolean = await queryInterface.tableExists('order_item')
		if (tableExist) {
			return queryInterface.dropTable('order_item', { logging: true })
		}
	}
}
