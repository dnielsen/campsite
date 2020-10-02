import {
  BaseEntity,
  Entity,
  PrimaryGeneratedColumn,
  ManyToMany,
  OneToMany,
  JoinTable,
  Column,
} from "typeorm";
import Session from "./Session";

@Entity("speakers")
export default class Speaker extends BaseEntity {
  @PrimaryGeneratedColumn("uuid")
  id!: string;

  @Column()
  name!: string;

  @Column()
  bio!: string;

  @Column()
  photo!: string;

  @Column()
  headline!: string;

  @ManyToMany(() => Session, (session) => session.speakers)
  @JoinTable({
    name: "session_speakers",
    joinColumn: {
      name: "speaker_id",
      referencedColumnName: "id",
    },
    inverseJoinColumn: {
      name: "session_id",
      referencedColumnName: "id",
    },
  })
  sessions!: Session[];
}
