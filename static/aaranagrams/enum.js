import { globalEn } from '../globalEnum.js';

var en = {
    Phase: {
        PREGAME: '0',
        PLAY: '1',
    },
    ToServerCode: {
        MIN_WORD_LENGTH: '2',
        SCORE_TO_WIN: '3',
        CHAOS_MODE: '4',
        START_GAME: 'a',
        END_GAME: 'b',
        GAME_MESSAGE: 'c',
        VOTE_SKIP: 'd',
        WHAT: 'e',
        WORD_ATTEMPT: 'f',
        WHAT: 'g',
        LETTER: 'h',
    },
    ToClientCode: {
        START_GAME: '0',
        END_GAME: '2',
        PLAYERS: '3',
        MIN_WORD_LENGTH: '5',
        SCORE_TO_WIN: '6',
        CHAOS_MODE: '7',
        GAME_MESSAGE: 'a',
        WHAT_RESPONSE: 'b',
        MESSAGE_WITH_WHAT: 'c',
        PROMPT: 'd',
        LETTERS: 'e',
        WORD_SUCCESS: 'f',
        YOUR_TURN: 'g'
    }
}
Object.assign(en.ToServerCode, globalEn.ToServerCode);
Object.assign(en.ToClientCode, globalEn.ToClientCode);
export { en };