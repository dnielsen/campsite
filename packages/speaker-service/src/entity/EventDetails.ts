import {
  BaseEntity,
  Entity,
  PrimaryGeneratedColumn,
  OneToMany,
  Column,
} from "typeorm";
import Session from "./Session";

@Entity("events")
export default class EventDetails extends BaseEntity {
  @PrimaryGeneratedColumn("uuid")
  id!: string;

  @Column()
  name!: string;

  @Column()
  description!: string;

  @Column()
  registrationUrl!: string;

  @Column()
  organizerName!: string;

  @Column()
  photo!: string;

  @Column()
  address!: string;

  @Column("timestamp", { name: "start_date" })
  startDate!: Date;

  @Column("timestamp", { name: "end_date" })
  endDate!: Date;

  @OneToMany(() => Session, (session) => session.event)
  sessions!: Session[];
}
