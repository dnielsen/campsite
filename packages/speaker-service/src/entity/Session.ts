import {
  BaseEntity,
  Entity,
  PrimaryGeneratedColumn,
  ManyToMany,
  Column,
  ManyToOne,
  JoinColumn,
} from "typeorm";
import Speaker from "./Speaker";
import EventDetails from "./EventDetails";

@Entity("sessions")
export default class Session extends BaseEntity {
  @PrimaryGeneratedColumn("uuid")
  id!: string;

  @Column()
  name!: string;

  @Column()
  description!: string;

  @Column()
  url!: string;

  @Column("timestamp", { name: "start_date" })
  startDate!: Date;

  @Column("timestamp", { name: "end_date" })
  endDate!: Date;

  @ManyToMany(() => Speaker, (speaker) => speaker.sessions)
  speakers!: Speaker[];

  @ManyToOne(() => EventDetails, (eventDetails) => eventDetails.sessions)
  @JoinColumn({ name: "event_id" })
  event!: EventDetails;
}
